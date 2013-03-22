package main

import (
	"log"
	"path/filepath"
	"os"
	"flag"
	"fmt"
)

// it's alive!
func main() {
	allPaths := []string{"abc", "def"}
	var scan = func(path string, _ os.FileInfo, inpErr error) (err error) {
		fullpath, err := filepath.Abs(path)
		if err != nil {
			return nil
		}
		allPaths = append(allPaths, fullpath)
		return nil
	}
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, scan)
	if err != nil {
		log.Fatal("Error walking directory")
	}
	for _, path := range allPaths {
		fmt.Printf("Scanned %s\n", path)
	}
}
