package main

import (
	"flag"
	"fmt"
	"os"
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
		fmt.Println("File size in bytes:", fileSize)
		fmt.Println("File Last Modified:", fileSummary.ModTime())

	// For "dir" case
	case "dir":
		dirname := os.Args[2]
		dirCmd.Parse(os.Args[3:])
		fmt.Printf("Directory Name: %s", dirname)
		fmt.Println("  tail:", dirCmd.Args())
		if *dirRecursive {
			fmt.Println("Summarizing Recursively")
		} else {
			fmt.Printf("Summarizing %s but not contained subfolders", dirname)
		}
	}
	if *verbose {
		fmt.Println("Verbose Summary Selected!")
	}
}