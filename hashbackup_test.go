package main

import (
	"crypto/md5"
	"crypto/sha1"
	"path"
	"runtime"
	"testing"
)

func sampleDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "test_data")
}

func sampleFile() string {
	return path.Join(sampleDir(),"foo.txt")
}

func TestGetPathsOfInterest(t *testing.T) {
	results, _ := getPathsOfInterest(sampleDir())
	if len(results) != 1 {
		t.Error("more than 1 item returned")
	}
	if results[0].path != sampleFile() {
		t.Error("incorrect file path returned")
	}
	if results[0].bytes != 20 {
		t.Error("incorrect file byte count returned")
	}
}

func TestCalculateAllHashes(t *testing.T) {
	input := fileDataSlice{}
	data := fileData{path: sampleFile()}
	input = append(input, data)
	results := calculateAllHashes(input)
	if len(results) != 1 {
		t.Error("more than 1 item returned")
	}
	if results[0].md5 == "" {
		t.Error("md5 must be set")
	}
	if results[0].sha1 == "" {
		t.Error("sha1 must be set")
	}
}

func TestGenHashWithMd5(t *testing.T) {
	result := genHash(md5.New(), sampleFile())
	if result != "4221d002ceb5d3c9e9137e495ceaa647" {
		t.Error("md5 not correct")
	}
}

func TestGenHashWithSha1(t *testing.T) {
	result := genHash(sha1.New(), sampleFile())
	if result != "804d716fc5844f1cc5516c8f0be7a480517fdea2" {
		t.Error("sha1 not correct")
	}
}

func TestCalculateHashes(t *testing.T) {
	fileData := fileData{path: sampleFile()}
	result := calculateHashes(fileData)
	if result.md5 == "" {
		t.Error("md5 must be set")
	}
	if result.sha1 == "" {
		t.Error("sha1 must be set")
	}
}
