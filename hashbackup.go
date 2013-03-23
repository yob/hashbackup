package main

import (
	"io"
	"log"
	"crypto/md5"
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

func hashPaths(allPaths []string) (results []string) {
	ch := make(chan string)
	for _, path := range allPaths {
		go func(path string) {
			ch <- hashPath(path)
		}(path)
	}
	for len(results) < len(allPaths) {
		result := <-ch
		results = append(results, result)
	}
	return results
}

func hashPath(path string) string {
	 file, err := os.Open(path)
	 if err != nil {
		return path
	 }
	 h := md5.New()
	 io.Copy(h, file)
	 file.Close()
	return path + " " + fmt.Sprintf("%x", h.Sum(nil))
}

// it's alive!
func main() {
	flag.Parse()
	root := flag.Arg(0)
	allPaths, err := getPathsOfInterest(root)
	if err != nil {
		log.Fatal("Error walking directory")
	}
	results := hashPaths(allPaths)
	for _, item := range results {
		fmt.Printf("hash: %s\n", item)
	}
}
