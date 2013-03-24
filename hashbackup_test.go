package main

import (
	"crypto/md5"
	"crypto/sha1"
	"path"
	"runtime"
	"testing"
)

func sampleFile() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "test_data","foo.txt")
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
