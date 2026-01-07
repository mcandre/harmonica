//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/harmonica"
	mageextras "github.com/mcandre/mage-extras"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// GeneratedTestFilePattern matches generated test files.
var GeneratedTestFilePattern = regexp.MustCompile("issue-")

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// Audit runs a security audit.
func Audit() error {return Govulncheck() }

// Test executes the integration test suite.
func Test() error {
	mg.Deps(Install)
	cmd := exec.Command("harmonica", "-version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Deadcode("./...") }

// DockerBuild creates local Docker buildx images.
func DockerBuild() error {
	return mageextras.Tuggy(
		"-t", fmt.Sprintf("n4jm4/harmonica:%s", harmonica.Version),
		"--load",
	)
}

// DockerPush creates and tag aliases remote Docker buildx images.
func DockerPush() error {
	return mageextras.Tuggy(
		"-t", fmt.Sprintf("n4jm4/harmonica:%s", harmonica.Version),
		"-a", "n4jm4/harmonica",
		"--push",
	)
}

// DockerTest creates and tag aliases remote test Docker buildx images.
func DockerTest() error {
	return mageextras.Tuggy(
		"-t", "n4jm4/harmonica:test",
		"--load",
		"--push",
	)
}

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck("./...") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("harmonica-%s", harmonica.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/harmonica"

// Factorio cross-compiles Go binaries for a multitude of platforms.
func Factorio() error {
	os.Setenv("FACTORIO_PLATFORM_BLOCKLIST", `(android\/.*)|(ios\/.*)|(.*\/wasm)|(plan9\/.*)`)
	return mageextras.Factorio(portBasename)
}

// Port builds and compresses artifacts.
func Port() error {
	mg.Deps(Factorio);

	return mageextras.Chandler(
		"-C",
		artifactsPath,
		"-czf",
		fmt.Sprintf("%s.tgz", portBasename),
		portBasename,
	)
}

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("harmonica") }

// Clean deletes junk files.
func Clean() error {
	err := CleanExamples()
	err2 := CleanArtifacts()

	if err != nil {
		return err
	}

	return err2
}

// CleanExamples deletes common generated test files.
func CleanExamples() error {
	examples := "examples"

	entries, err := os.ReadDir(examples)

	if err != nil {
		return err
	}

	for _, entry := range entries {
		pth := filepath.Join(examples, entry.Name())

		if !GeneratedTestFilePattern.MatchString(pth) {
			continue
		}

		fi, err2 := os.Stat(pth)

		if err2 != nil {
			return err2
		}

		if fi.IsDir() {
			if err3 := os.RemoveAll(pth); err3 != nil {
				return err3
			}
		} else {
			if err3 := os.Remove(pth); err3 != nil {
				return err3
			}
		}
	}

	return nil
}

// CleanArtifacts deletes application ports.
func CleanArtifacts() error { return os.RemoveAll(artifactsPath) }
