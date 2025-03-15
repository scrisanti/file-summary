package main

import (
	"os"
)

func fileDescribe(fn string) (os.FileInfo, error) {
	// Determine File Type and size
	fileInfo, err := os.Stat(fn)
	// Read File and Describe
	return fileInfo, err
}