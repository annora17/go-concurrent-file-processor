package main

import (
	"fmt"
	"processor/processor"
)

func main() {
	targetDirectory := "testdata"

	fileProcessor := processor.NewFileProcessor(targetDirectory, "")

	result := fileProcessor.ProcessFile("AT")

	for _, v := range result.FileSearch {
		fmt.Printf("filename -> %s || matched count -> %d\n", v.FileName, v.MatchCount)

	}
}
