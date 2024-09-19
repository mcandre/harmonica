package harmonica

import (
	"github.com/saracen/fastzip"

	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Copy file paths.
func Copy(dest string, source string) error {
	fIn, err := os.Open(source)

	defer func() {
		if err2 := fIn.Close(); err2 != nil {
			log.Println(err2)
		}
	}()

	if err != nil {
		return err
	}

	fOut, err := os.Create(dest)

	if err != nil {
		return err
	}

	defer func() {
		if err2 := fOut.Close(); err2 != nil {
			log.Print(err2)
		}
	}()

	_, err = io.Copy(fOut, fIn)
	return err
}

// GetDirectorySizeBytes totals directory content file size recursively,
// measured in bytes.
func GetDirectorySizeBytes(root string) (int64, error) {
	var sum int64

	err := filepath.Walk(root, func(pth string, fi fs.FileInfo, err2 error) error {
		if err2 != nil {
			return err2
		}

		if !fi.IsDir() {
			sum += fi.Size()
		}

		return nil
	})

	return sum, err
}

// Config models a Harmonica operation.
type Config struct {
	// Source denotes a file path (required, nonblank).
	Source string

	// Prefix denotes a target basename.
	Prefix string

	// ExpandSource denotes whether to treat the source initially as a ZIP format archive.
	// If so, attempt to expand the archive.
	// Assume that the archive follows a convention
	// of nesting all of its entries in an envelope directory,
	// with the directory name identical to the archive path minus the archive file extension.
	ExpandSource bool

	// Assets denotes a collection of boilerplate asset files to copy into each batch.
	Assets []string

	// BatchLimitEntries denotes a cap on the number of files in each batch.
	// 0 indicates unbounded.
	BatchLimitEntries uint

	// BatchLimitMiB denotes a cap on the (uncompressed) content size in each batch.
	// 0 indicates unbounded.
	BatchLimitMiB uint

	// BatchCompressionExtension denotes an optional ZIP format file extension,
	// E.g. "zip", "cbz", "jar".
	// Blank indicates no compression.
	BatchCompressionExtension string

	// cwdAbs denotes the absolute current working directory.
	cwdAbs string

	// sourceExpandedAbs denotes the absolute source directory,
	// potentially after expanding a source archive.
	sourceExpandedAbs string

	// prefixAbs denotes the absolute target directory prefix,
	// potentially relative to the current working directory's parent.
	prefixAbs string

	// sources collects the original files to distribute.
	sources []string

	// batchID tracks the current batch identifier.
	batchID uint

	// fileID tracks the current file identifier.
	fileID uint

	// targetAbs tracks the absolute path to the current batch target.
	targetAbs string
}

// Run launches a Harmonica operation.
func (o *Config) Run() error {
	if o.Source == "" {
		return fmt.Errorf("blank source")
	}

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	o.cwdAbs, err = filepath.Abs(cwd)

	if err != nil {
		return err
	}

	sourceAbs, err := filepath.Abs(o.Source)

	if err != nil {
		return nil
	}

	if o.ExpandSource {
		extractor, err2 := fastzip.NewExtractor(sourceAbs, o.cwdAbs)

		if err2 != nil {
			return err2
		}

		defer func() {
			if err3 := extractor.Close(); err2 != nil {
				log.Println(err3)
			}
		}()

		if err3 := extractor.Extract(context.Background()); err3 != nil {
			return err3
		}

		o.sourceExpandedAbs, err2 = filepath.Abs(
			strings.TrimSuffix(sourceAbs, filepath.Ext(sourceAbs)),
		)

		if err2 != nil {
			return err2
		}

	} else {
		o.sourceExpandedAbs = sourceAbs
	}

	fi, err := os.Stat(o.sourceExpandedAbs)

	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("unable to read directory: %s", o.sourceExpandedAbs)
	}

	if o.sourceExpandedAbs == o.cwdAbs {
		log.Printf("warning: relocating batches to parent directory for safety")
		o.prefixAbs = filepath.Join(filepath.Dir(o.cwdAbs), o.Prefix)
	} else {
		o.prefixAbs = filepath.Join(o.cwdAbs, o.Prefix)
	}

	if err2 := filepath.WalkDir(o.sourceExpandedAbs, o.QueuePath); err2 != nil {
		return err2
	}

	if len(o.sources) == 0 {
		return fmt.Errorf("empty source: %s", o.sourceExpandedAbs)
	}

	if err2 := o.nextBatch(); err2 != nil {
		return err2
	}

	for _, source := range o.sources {
		beyondCapacity, err3 := o.BeyondBatchCapacity()

		if err3 != nil {
			return err3
		}

		if beyondCapacity {
			if err4 := o.nextBatch(); err4 != nil {
				return err4
			}
		}

		target := filepath.Join(o.targetAbs, filepath.Base(source))

		if err4 := Copy(target, source); err4 != nil {
			return err4
		}

		o.fileID++
	}

	return o.finalizeBatch()
}

// BeyondBatchCapacity reports whether the current batch is too full
// to receive additional files.
func (o Config) BeyondBatchCapacity() (bool, error) {
	if o.BatchLimitEntries > 0 && o.fileID >= o.BatchLimitEntries {
		return true, nil
	}

	if o.BatchLimitMiB > 0 {
		b, err := GetDirectorySizeBytes(o.targetAbs)

		if err != nil {
			return false, err
		}

		mib := uint(b / 1048576)

		if mib >= o.BatchLimitMiB {
			return true, nil
		}
	}

	return false, nil
}

// QueuePath collects non-directory source paths.
func (o *Config) QueuePath(pth string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if !d.IsDir() {
		o.sources = append(o.sources, pth)
	}

	return nil
}

func (o *Config) nextBatch() error {
	if err := o.finalizeBatch(); err != nil {
		return err
	}

	o.batchID++
	o.targetAbs = fmt.Sprintf("%s%d", o.prefixAbs, o.batchID)

	if err := os.MkdirAll(o.targetAbs, 0755); err != nil {
		return err
	}

	for _, asset := range o.Assets {
		target := filepath.Join(o.targetAbs, filepath.Base(asset))

		if err := Copy(target, asset); err != nil {
			return err
		}
	}

	o.fileID = uint(len(o.Assets))
	return nil
}

// finalizeBatch optionally compresses the current batch.
func (o *Config) finalizeBatch() error {
	if o.BatchCompressionExtension == "" ||
		o.batchID == 0 {
		return nil
	}

	targetArchive := fmt.Sprintf("%s%s", o.targetAbs, o.BatchCompressionExtension)
	archiveWriter, err := os.Create(targetArchive)

	if err != nil {
		return err
	}

	defer func() {
		if err2 := archiveWriter.Close(); err2 != nil {
			log.Println(err2)
		}
	}()

	archiver, err := fastzip.NewArchiver(archiveWriter, o.targetAbs)

	if err != nil {
		return err
	}

	defer func() {
		if err2 := archiver.Close(); err2 != nil {
			log.Println(err2)
		}
	}()

	files := make(map[string]os.FileInfo)

	err = filepath.Walk(o.targetAbs, func(pth string, fi fs.FileInfo, err2 error) error {
		if err2 != nil {
			return err2
		}

		files[pth] = fi
		return nil
	})

	if err != nil {
		return nil
	}

	return archiver.Archive(context.Background(), files)
}
