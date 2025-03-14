package main

import (
	"os"
	"flag"
	"fmt"
)

func main() {
	
	// Flags for single files
	fileCmd			:= flag.NewFlagSet("file", flag.ExitOnError)
	// fileName 		:= fileCmd.String("filename", "", "Filename to analyze")

	// Flags for directories 
	dirCmd			:= flag.NewFlagSet("dir", flag.ExitOnError)
	// dirName 		:= dirCmd.String("directory", "", "Directory to Analyze")
	dirRecursive 	:= dirCmd.Bool("recusrive",false,"Should all sub-directories be analyzed?")

	// Define Global Flags
	verbose 		:= flag.Bool("verbose", false, "Enable Verbose Output")

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
		// fmt.Printf("Analyzing File %s", *fileName)
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