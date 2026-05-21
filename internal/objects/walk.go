package objects

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WalkTree(path string) (string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}
	var builder strings.Builder

	for _, entry := range entries {
		if entry.Name() == ".goat" || entry.Name() == ".git" {
			continue
		}
		if !entry.IsDir() {
			blobHash, err := Store(filepath.Join(path, entry.Name()))
			if err != nil {
				return "", err
			}
			builder.WriteString(fmt.Sprintf("blob %s %s\n", entry.Name(), blobHash))
		} else {
			treeHash, err := WalkTree(filepath.Join(path, entry.Name()))
			if err != nil {
				return "", err
			}
			builder.WriteString(fmt.Sprintf("tree %s %s\n", entry.Name(), treeHash))
		}
	}
	entriesString := builder.String()
	treeObject := fmt.Sprintf("tree %d\x00%s", len(entriesString), entriesString)
	fmt.Println(treeObject)
	treeHash, err := StoreObject(treeObject)
	if err != nil {
		return "", err
	}
	return treeHash, nil
}
