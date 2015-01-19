package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/cover"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Syntax: gocov-merge FILENAME\n")
		os.Exit(1)
	}
	profiles, err := cover.ParseProfiles(os.Args[1])
	if err != nil {
		fmt.Println("Unable to parse file")
		os.Exit(1)
	}
	fmt.Print("mode: count\n")
	for _, profile := range profiles {
		p := &Profile{profile, make([]*ProfileBlock, 0)}
		p.MergeBlocks()
		fmt.Print(p.Format())
	}
}
