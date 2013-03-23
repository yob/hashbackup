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
	"sort"
)

type fileData struct {
	path  string
	md5   string
	sha1  string
	bytes int64
}

type fileDataSlice []fileData

func (s fileDataSlice) Less(i, j int) bool {
	return s[i].path < s[j].path
}

func (s fileDataSlice) Len() int {
	return len(s)
}

func (s fileDataSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// returns a slice of file paths the user is interested in. root is the top
// level directory to search under
func getPathsOfInterest(root string) (allData fileDataSlice, err error) {
	var scan = func(path string, info os.FileInfo, _ error) error {
		if !info.IsDir() {
			fullpath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			data := fileData{path: fullpath, bytes: info.Size()}
			allData = append(allData, data)
		}
		return nil
	}
	err = filepath.Walk(root, scan)
	return
}

func calculateAllHashes(allData fileDataSlice) (results fileDataSlice) {
	ch := make(chan fileData)
	for _, path := range allData {
		go func(data fileData) {
			ch <- calculateHashes(data)
		}(path)
	}
	for len(results) < len(allData) {
		result := <-ch
		results = append(results, result)
	}
	sort.Sort(results)
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

func calculateHashes(info fileData) (fileData) {
	info.md5  = genHash(md5.New(), info.path)
	info.sha1 = genHash(sha1.New(), info.path)
	return info
}

// it's alive!
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	root := flag.Arg(0)
	allData, err := getPathsOfInterest(root)
	if err != nil {
		log.Fatal("Error walking directory")
	}
	results := calculateAllHashes(allData)
	for _, data := range results {
		fmt.Printf("%s %s %d %s\n", data.md5, data.sha1, data.bytes, data.path)
	}
}
