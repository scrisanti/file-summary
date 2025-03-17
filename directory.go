package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
	// "sync"
	//"github.com/h2non/filetype"
)

func mainDirAnalyze(rootDir string) ([]string, int, error){
	files,totalFileSize, err := walkDirectory(rootDir)
	
	var wg sync.WaitGroup
	runningSum := 0

	timeStart := time.Now()
	for _, file := range files { 
		wg.Add(1)

		go func() {
			defer wg.Done()
			lineCount, err := worker(file)
				if err != nil {
					log.Fatal(err)
				}
			fmt.Printf("%s has %d lines\n", filepath.Base(file), lineCount)
			runningSum += lineCount

		}()
		
		}
	
	wg.Wait()
	timeEnd := time.Now()
	fmt.Printf("Total Time: %v\n", timeEnd.Sub(timeStart))
	fmt.Printf("Total Lines: %d\n", runningSum)

	return files,totalFileSize, err
}


func walkDirectory(rootDir string) ([]string, int, error) {
	var files []string
	var totalFileSize int
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		// TODO: Implement Recursive File Description
		// Only look at sub-directories if recursive is true
		if !info.IsDir() {
			files = append(files, path)
			totalFileSize += int(info.Size())
		}
		return nil
	})
	return files,totalFileSize, err
}

func worker(fn string) (int, error){
	fmt.Printf("Reading File: %s\n", filepath.Base(fn))
	fileStartRead := time.Now()

	file, err := os.Open(fn)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Some files need larger buffers
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err	
	}
	fileEndRead := time.Now()
	fmt.Printf("Done Reading: %s - %v\n", filepath.Base(fn), fileEndRead.Sub(fileStartRead))

	return lineCount, nil

}