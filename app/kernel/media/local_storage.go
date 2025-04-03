package media

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	BasePath string // Folder where files will be saved.
	BaseURL  string // URL prefix to access the files.
}

func (ls *LocalStorage) Save(file multipart.File, filename string) (string, error) {
	// Ensure the storage directory exists.
	if err := os.MkdirAll(ls.BasePath, os.ModePerm); err != nil {
		return "", fmt.Errorf("creating storage directory: %w", err)
	}

	// Create the destination file.
	dstPath := filepath.Join(ls.BasePath, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("creating file: %w", err)
	}

	defer dst.Close()

	// Copy file contents.
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("copying file: %w", err)
	}

	// Construct the file URL.
	return fmt.Sprintf("%s/%s", ls.BaseURL, filename), nil
}
