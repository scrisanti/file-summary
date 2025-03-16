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
	// Go Routine for meta analysis
	// ch := make(chan int) // Create a channel of type int
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

	file, err := os.Open(fn)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err	
	}
	fmt.Printf("Done Reading: %s\n", filepath.Base(fn))

	return lineCount, nil

}