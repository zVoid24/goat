package main

import (
	"fmt"
	"goat/internal/objects"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func walk(path string, depth int) (int, int) {
	entries, err := os.ReadDir(path)
	fileCount := 0
	dirCount := 0
	if err != nil {
		fmt.Println(err)
		return fileCount, dirCount
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		label := "[DIR]"
		if !entry.IsDir() {
			label = "[FILE]"
			fileCount++
		} else {
			dirCount++
		}
		fmt.Println(strings.Repeat("  ", depth) + label + " " + filepath.Join(path, entry.Name()))
		if entry.IsDir() {
			childFile, childDir := walk(filepath.Join(path, entry.Name()), depth+1)
			fileCount += childFile
			dirCount += childDir
		}
	}
	return fileCount, dirCount
}

func main() {
	// fmt.Println("hello world")
	// fileCount, dirCount := walk(".", 0)
	// fmt.Println("Total Files", fileCount)
	// fmt.Println("Total Directories", dirCount)
	if os.Args[1] == "cat" {
		content, err := objects.Read(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(content)
		return
	}
	hash, err := objects.Store(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hash)
}
