package main

import (
	"github.com/mcandre/harmonica"

	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var flagPrefix = flag.String("prefix", "issue-", "")
var flagN = flag.Int("n", 0, "Cap batch entry count")
var flagM = flag.Int("m", 250, "Cap batch directory size (MiB)")
var flagUnzip = flag.Bool("unzip", false, "Pre-process expand ZIP/CBZ source archive")
var flagAssets = flag.String("assets", "", "Copy (comma separated) asset files to each batch")
var flagZip = flag.String("zip", "", "ZIP format batch compression file extension (e.g. \".zip\", \".cbz\", \".jar\")")
var flagVersion = flag.Bool("version", false, "Show version")
var flagHelp = flag.Bool("help", false, "Show usage menu")

// usage presents a help menu.
func usage() {
	fmt.Printf("Usage: %s [OPTIONS] <source>\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	switch {
	case *flagVersion:
		fmt.Printf("%v\n", harmonica.Version)
		os.Exit(0)
	case *flagHelp:
		usage()
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) != 1 {
		usage()
		os.Exit(1)
	}

	var assets []string

	assetsCommaSeparated := *flagAssets

	if assetsCommaSeparated != "" {
		assets = strings.Split(assetsCommaSeparated, ",")
	}

	config := harmonica.Config{
		Source:                    args[0],
		Prefix:                    *flagPrefix,
		ExpandSource:              *flagUnzip,
		Assets:                    assets,
		BatchLimitEntries:         uint(*flagN),
		BatchLimitMiB:             uint(*flagM),
		BatchCompressionExtension: *flagZip,
	}

	if err := config.Run(); err != nil {
		log.Fatal(err)
	}
}
