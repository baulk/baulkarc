package main

import (
	"archive/zip"
	"fmt"
	"os"
)

// trace zip details

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s zipfile\n", os.Args[0])
		os.Exit(1)
	}
	reader, err := zip.OpenReader(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable open zipfile: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()
	for _, entry := range reader.File {
		if entry.Flags&0x800 != 0 {
			fmt.Fprintf(os.Stderr, "%s (UTF-8)\n", entry.Name)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", entry.Name)
		}
	}
}
