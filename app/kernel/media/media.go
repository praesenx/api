package media

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

type Media struct {
	file       []byte
	headerName string
	ext        string
	filename   string
	filePath   string
}

func MakeMedia(file []byte, headerName string) (*Media, error) {
	if len(file) <= 0 {
		return nil, errors.New("the given file is empty")
	}

	ext := filepath.Ext(headerName)
	filename := uuid.New().String() + ext

	media := Media{
		file:       file,
		headerName: headerName,
		ext:        ext,
		filename:   filename,
		filePath:   filepath.Join(GetUsersImagesDir(), filename),
	}

	return &media, nil
}

func (m Media) Write() error {
	err := os.WriteFile(m.filePath, m.file, 0644)

	if err != nil {
		return err
	}

	return nil
}
