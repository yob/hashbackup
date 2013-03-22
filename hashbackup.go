package main

import (
	"log"
	"path/filepath"
	"os"
	"flag"
	"fmt"
)

// returns a slice of file paths the user is interested in. root is the top
// level directory to search under
func getPathsOfInterest(root string) (allPaths []string, err error) {
	var scan = func(path string, info os.FileInfo, _ error) error {
		if !info.IsDir() {
			fullpath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			allPaths = append(allPaths, fullpath)
		}
		return nil
	}
	err = filepath.Walk(root, scan)
	return
}

// it's alive!
func main() {
	flag.Parse()
	allPaths, err := getPathsOfInterest(flag.Arg(0))
	if err != nil {
		log.Fatal("Error walking directory")
	}
	for _, path := range allPaths {
		fmt.Printf("Scanned %s\n", path)
	}
}
