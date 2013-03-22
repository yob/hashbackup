package main

import (
	"path/filepath"
	"os"
	"flag"
	"fmt"
)

// returns a slice of file paths the user is interested in. root is the top
// level directory to search under
func getPathsOfInterest(root string) []string {
	allPaths := []string{}
	var scan = func(path string, _ os.FileInfo, inpErr error) (err error) {
		fullpath, err := filepath.Abs(path)
		if err != nil {
			return nil
		}
		allPaths = append(allPaths, fullpath)
		return nil
	}
	filepath.Walk(root, scan)
	return allPaths
}

// it's alive!
func main() {
	flag.Parse()
	allPaths := getPathsOfInterest(flag.Arg(0))
	for _, path := range allPaths {
		fmt.Printf("Scanned %s\n", path)
	}
}
