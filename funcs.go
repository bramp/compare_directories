package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pmezard/go-difflib/difflib"
)

func hashFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func diff(a, b []string) []string {
	var onlyInA []string
	for _, x := range a {
		if !contains(b, x) {
			onlyInA = append(onlyInA, x)
		}
	}
	return onlyInA
}

func diffContent(a, b []string, dir1, dir2 string) []string {
	var diffContent []string
	for _, x := range a {
		if contains(b, x) {
			if !isFileIdentical(filepath.Join(dir1, x), filepath.Join(dir2, x)) {
				diffContent = append(diffContent, x)
				file1, err := ioutil.ReadFile(filepath.Join(dir1, x))
				if err != nil {
					fmt.Printf("Failed to read file %s: %v\n", x, err)
					continue
				}
				file2, err := ioutil.ReadFile(filepath.Join(dir2, x))
				if err != nil {
					fmt.Printf("Failed to read file %s: %v\n", x, err)
					continue
				}
				diffs := difflib.UnifiedDiff{
					A:        difflib.SplitLines(string(file1)),
					B:        difflib.SplitLines(string(file2)),
					FromFile: x,
					ToFile:   x,
					Context:  3,
				}
				text, _ := difflib.GetUnifiedDiffString(diffs)
				fmt.Println(text)
			}
		}
	}
	return diffContent
}

func contains(list []string, x string) bool {
	for _, y := range list {
		if x == y {
			return true
		}
	}
	return false
}

func isFileIdentical(file1, file2 string) bool {
	h1, err := hashFile(file1)
	if err != nil {
		return false
	}

	h2, err := hashFile(file2)
	if err != nil {
		return false
	}

	return bytes.Equal(h1, h2)
}
