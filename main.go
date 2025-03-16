package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	
	// Flags for single files
	fileCmd			:= flag.NewFlagSet("file", flag.ExitOnError)
	verbose 		:= fileCmd.Bool("verbose", false, "Enable Verbose Output")

	// Flags for directories 
	dirCmd			:= flag.NewFlagSet("dir", flag.ExitOnError)
	dirRecursive 	:= dirCmd.Bool("recusrive",false,"Should all sub-directories be analyzed?")

	// Subcommand must be first Command
	if len(os.Args) < 2 {
        fmt.Println("expected 'file' or 'dir' subcommands")
        os.Exit(1)
    }

	// ------- Check All Commands ----------- //
	switch os.Args[1] {
	
		// For "file" case
	case "file":
		filename := os.Args[2]
		fileCmd.Parse(os.Args[3:])
		fmt.Printf("Filename: %s", filename)
		fmt.Println("  tail:", fileCmd.Args())

		fileSummary, err := fileDescribe(filename)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fileSize := fileSummary.Size()
		fmt.Println("File Size in bytes:", fileSize)
		fmt.Println("File Last Modified:", fileSummary.ModTime())

		if *verbose {
			var wg sync.WaitGroup

			fmt.Println("Verbose Summary Selected!")
			if isImgFile(filename) {
				imgFiles := []string{}
				imgFiles = append(imgFiles, filename) // Only one file
				// Pass list of images to the verbose summarizer
				for _, fn := range imgFiles {
					wg.Add(1)
					go verboseImageSummary(fn, &wg) // verboseImageSummary(fn, &wg)
				}
				// Wait for all goroutines to finish
				wg.Wait()
				fmt.Println("Processing complete.")
			}
		}

	// For "dir" case
	case "dir":
		dirname := os.Args[2]
		dirCmd.Parse(os.Args[3:])
		fmt.Printf("Directory Name: %s", dirname)
		fmt.Println("  tail:", dirCmd.Args())
		if *dirRecursive {
			fmt.Println("Summarizing Recursively")
		} else {
			fmt.Printf("Summarizing %s but not contained subfolders\n", dirname)
			// Find all files in folder and group by type
			files, totalFileSize, err := mainDirAnalyze(dirname)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Found %d files - %d MB\n", len(files), totalFileSize/1000000)
		}
	}
	
}