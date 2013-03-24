package main

import (
	"crypto/md5"
	"crypto/sha1"
	"path"
	"runtime"
	"testing"
)

func TestGenHashWithMd5(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	f := path.Join(path.Dir(filename), "test_data","foo.txt")
	result := genHash(md5.New(), f)
	if result != "4221d002ceb5d3c9e9137e495ceaa647" {
		t.Log("md5 not correct")
		t.Fail()
	}
}

func TestGenHashWithSha1(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	f := path.Join(path.Dir(filename), "test_data","foo.txt")
	result := genHash(sha1.New(), f)
	if result != "804d716fc5844f1cc5516c8f0be7a480517fdea2" {
		t.Log("sha1 not correct")
		t.Fail()
	}
}
