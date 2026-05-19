package objects

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Store(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	header := fmt.Sprintf("blob %d\x00", len(fileContent))
	fileObject := append([]byte(header), fileContent...)
	var buf bytes.Buffer
	comprsr := zlib.NewWriter(&buf)
	comprsr.Write(
		fileObject,
	)
	comprsr.Close()
	fmt.Println(buf)
	fmt.Printf("%q\n", fileObject)

	h := sha1.New()
	if _, err := h.Write(fileObject); err != nil {
		return "", err
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))
	objDir := hash[:2]
	objFile := hash[2:]
	err = os.MkdirAll(filepath.Join(".goat/objects", objDir), 0755)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(filepath.Join(".goat/objects", objDir, objFile), buf.Bytes(), 0644)
	if err != nil {
		return "", err
	}
	return hash, nil
}
