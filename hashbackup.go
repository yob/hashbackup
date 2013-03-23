package main

import (
	"io"
	"log"
	"hash"
	"crypto/md5"
	"crypto/sha1"
	"path/filepath"
	"os"
	"flag"
	"fmt"
	"runtime"
)

type fileInfo struct {
	path  string
	md5   string
	sha1  string
	bytes int
}

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

func loadAllInfo(allPaths []string) (results []fileInfo) {
	ch := make(chan fileInfo)
	for _, path := range allPaths {
		go func(path string) {
			ch <- loadInfo(path)
		}(path)
	}
	for len(results) < len(allPaths) {
		result := <-ch
		results = append(results, result)
	}
	return results
}

func genHash(hash hash.Hash, path string) string {
	file, err := os.Open(path)
	if err != nil {
		return path
	}
	io.Copy(hash, file)
	file.Close()
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func loadInfo(path string) (info fileInfo) {
	info.path = path
	info.md5  = genHash(md5.New(), path)
	info.sha1 = genHash(sha1.New(), path)
	return
}

// it's alive!
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	root := flag.Arg(0)
	allPaths, err := getPathsOfInterest(root)
	if err != nil {
		log.Fatal("Error walking directory")
	}
	results := loadAllInfo(allPaths)
	for _, info := range results {
		fmt.Printf("%s %s %s\n", info.md5, info.sha1, info.path)
	}
}
