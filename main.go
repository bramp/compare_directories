package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// Define command-line flags
	var help bool
	var dir1, dir2 string

	flag.BoolVar(&help, "help", false, "show this help message (shorthand: -h)")
	flag.BoolVar(&help, "h", false, "show this help message (shorthand for --help)")

	flag.StringVar(&dir1, "d1", "", "the first directory to compare (shorthand for -dir1)")
	flag.StringVar(&dir1, "dir1", "", "the first directory to compare (shorthand: -d1)")

	flag.StringVar(&dir2, "d2", "", "the first directory to compare (shorthand for -dir2)")
	flag.StringVar(&dir2, "dir2", "", "the first directory to compare (shorthand: -d1)")

	flag.Parse()

	// Check if help flag is set
	if help {
		flag.Usage()
		os.Exit(0)
	}

	// Check if both dir flags are set
	if dir1 == "" || dir2 == "" {
		fmt.Println("Both --dir1 and --dir2 flags are required.")
		flag.Usage()
		os.Exit(1)
	}

	var files1 []string
	err := filepath.Walk(dir1, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			relpath, _ := filepath.Rel(dir1, path)
			files1 = append(files1, relpath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Failed to read directory %s: %v\n", dir1, err)
		os.Exit(1)
	}

	var files2 []string
	err = filepath.Walk(dir2, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			relpath, _ := filepath.Rel(dir2, path)
			files2 = append(files2, relpath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Failed to read directory %s: %v\n", dir2, err)
		os.Exit(1)
	}

	// compare the lists of files
	onlyIn1 := diff(files1, files2)
	onlyIn2 := diff(files2, files1)

	diffContent := diffContent(files1, files2, dir1, dir2)

	if len(onlyIn1) == 0 && len(onlyIn2) == 0 && len(diffContent) == 0 {
		fmt.Println("The two directories are identical.")
	} else {
		if len(onlyIn1) > 0 {
			fmt.Println("Files only in", dir1)
			for _, file := range onlyIn1 {
				fmt.Println("-", file)
			}
		}

		if len(onlyIn2) > 0 {
			fmt.Println("Files only in", dir2)
			for _, file := range onlyIn2 {
				fmt.Println("-", file)
			}
		}

		if len(diffContent) > 0 {
			fmt.Println("Identical files with different content in both directories")
			for _, file := range diffContent {
				fmt.Println("-", file)
			}
		}

	}
}
