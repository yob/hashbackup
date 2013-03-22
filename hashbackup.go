package main

import (
	"path/filepath"
	"os"
	"flag"
	"fmt"
)

func visit(path string, f os.FileInfo, err error) error {
  fullpath, err := filepath.Abs(path)
  if err != nil {
	return nil
  }
  fmt.Printf("Visited: %s\n", fullpath)
  return nil
}

// it's alive!
func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

