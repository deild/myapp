// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

const (
	binary = "myapp"
)

var (
	Default   = Myapp
	version   = "dev"
	buildDate = time.Now().UTC().Format(time.RFC3339)
)

// Build myapp binary
func Myapp() error {
	fmt.Println("+ build")
	return sh.Run("go", "build", ldflags(), "-o", binary, ".")
}

// Install myapp binary
func Install() error {
	mg.Deps(Vendor)
	fmt.Println("+ install")
	getVersionFromTag()
	return sh.Run("go", "install", ldflags(), ".")
}

func getDep() error {
	return sh.Run("go", "get", "-u", "github.com/golang/dep/cmd/dep")
}

// Install Go Dep and sync vendored dependencies
func Vendor() error {
	mg.Deps(getDep)
	fmt.Println("+ vendor")
	return sh.Run("dep", "ensure")
}

// Formats the code via gofmt
func Fmt() error {
	fmt.Println("+ formatting code")
	fileList, err := goFileList()
	if err != nil {
		return err
	}
	for _, file := range fileList {
		err = sh.Run("gofmt", "-s", "-l", "-w", file)
		if err != nil {
			return err
		}
	}
	return nil
}

// Run all the tests
func Test() error {
	fmt.Println("+ test")
	return sh.Run("go", "test", "-covermode=atomic", "-coverprofile=coverage.out", "./...")
}

func getLint() error {
	return sh.Run("go", "get", "-u", "github.com/golang/lint/golint")
}

// Run Go Meta Linter
func Lint() error {
	mg.Deps(getLint)
	fmt.Println("+ lint")
	return sh.Run("golint", "./...")
}

// Run all the tests and code checks
func Ci() {
	mg.Deps(Clean)
	fmt.Println("+ ci")
	mg.Deps(Vendor)
	mg.Deps(Myapp)
	mg.Deps(Test, Fmt, Lint)
}

// Removes the generated directories and files
func Clean() {
	fmt.Println("+ clean")
	os.RemoveAll("vendor")
	os.RemoveAll(binary)
	os.RemoveAll("coverage.out")
	os.RemoveAll("dist")
	os.Remove(filepath.Join(os.Getenv("GOPATH"), "bin", binary))
}

// Create a snapshot with goreleaser
func Snapshot() error {
	return sh.Run("goreleaser", "--rm-dist", "--snapshot")
}

func ldflags() string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return fmt.Sprintf("-ldflags=-s -w -X main.date=%s -X main.commit=%s -X main.version=%s", buildDate, hash, version)
}

func goFileList() ([]string, error) {
	fileList := make([]string, 0)
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		if !strings.HasPrefix(path, "vendor") && strings.HasSuffix(path, ".go") {
			fileList = append(fileList, path)
		}
		return err
	})
	return fileList, err
}

func getVersionFromTag() {
	hashtag, _ := sh.Output("git", "rev-list", "--tags", "--max-count=1")
	if hashtag != "" {
		tag, _ := sh.Output("git", "describe", "--tags", hashtag)
		if tag != "" {
			version = tag
		}
	}
}
