package media

import (
	"errors"
	"github.com/google/uuid"
	"path/filepath"
)

const Dir = "media"
const UsersDir = "users"
const PostsDir = "posts"
const StorageDir = "storage"

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
	}

	media.filePath = filepath.Join(
		media.GetUsersImagesDir(),
		filename,
	)

	return &media, nil
}
