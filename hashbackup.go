package main

import (
	"crypto/md5"
	"crypto/sha1"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

// a convenient place to hold data on a signle file
type fileData struct {
	path  string
	md5   string
	sha1  string
	bytes int64
}

// a slice for holding a collection of fileData structs. Includes extra methods
// required to match sort.Interface so the collection is sortable.
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

// returns a slice of fileData structs the user is interested in. The structs
// have basic info (path, bytes) filled in but the hash fields are empty.
// root is the top level directory to search under
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
	sort.Sort(allData)
	return
}

// takes a collection of fileData structs, fills in the hash fields and returns
// a new slice with the complete data. Farms the work out to a go-routine per
// file for maximum concurrency.
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

// calculates a hash (md5, sha1, etc) for a single file. The calling function
// is responsible for providing the hash.Hash implementation that will be used.
func genHash(hash hash.Hash, path string) string {
	file, err := os.Open(path)
	if err != nil {
		return path
	}
	io.Copy(hash, file)
	file.Close()
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// fill in the hash values on the provided fileData struct.
func calculateHashes(info fileData) fileData {
	info.md5 = genHash(md5.New(), info.path)
	info.sha1 = genHash(sha1.New(), info.path)
	return info
}

// application entry point
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
		fmt.Printf("%s\t%s\t%d\t%s\n", data.md5, data.sha1, data.bytes, data.path)
	}
}
