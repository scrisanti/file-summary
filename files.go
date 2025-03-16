package main

import (
	"os"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
	"log"
	"fmt"
	//"io"
	"sync"
	"bufio"

	"github.com/h2non/filetype"
)

func fileDescribe(fn string) (os.FileInfo, error) {
	// Determine File Type and size
	fileInfo, err := os.Stat(fn)
	// Read File and Describe
	return fileInfo, err
}

func isImgFile(fn string) bool {
	// 	Open File"
	file, err := os.Open(fn)
	if err != nil {
		log.Fatalf("Could Not Open File: %v", err)
	}
	defer file.Close()

	head := make([]byte, 261) // Need min 261 bytes to determine filetype
	_, err = file.Read(head)
	if err != nil {
		log.Fatalf("Could Not Open File: %v", err)
	}
	// Check if image type 
	return filetype.IsImage(head) 
}

func verboseImageSummary(fn string, wg *sync.WaitGroup) {
	defer wg.Done() // Hold off so that all files process before exiting goroutine

	// 	Open File
	file, err := os.Open(fn)

	if err != nil {
		log.Fatalf("Could Not Open File: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		fmt.Printf("Error decoding image %s: %v\n", fn, err)
		return
	}

	fmt.Printf("%s: %dx%d pixels\n", filepath.Base(fn), img.Width, img.Height)

}