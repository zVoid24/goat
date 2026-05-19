package objects

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
	"path/filepath"
)

func Read(hash string) (string, error) {
	objDir := hash[:2]
	objFile := hash[2:]
	reqFile, err := os.ReadFile(filepath.Join(".goat/objects", objDir, objFile))
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(reqFile)
	reader, err := zlib.NewReader(buf)
	if err != nil {
		return "", err
	}
	defer reader.Close()
	content, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	_, realContent, _ := bytes.Cut(content, []byte("\x00"))
	return string(realContent), nil
}
