package media

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"path/filepath"
	"strings"
)

func MakeMedia(uniqueId string, file []byte, baseFileName string) (*Media, error) {
	if len(file) <= 0 {
		return nil, errors.New("the given file is empty")
	}

	fmt.Println("---> ", baseFileName)
	size := int64(len(file))
	ext := strings.ToLower(filepath.Ext(baseFileName))
	filename := uniqueId + "-" + uuid.New().String() + ext

	if size > maxFileSize {
		return nil, errors.New("the given file is too big")
	}

	if hasInvalidExt(ext) {
		return nil, errors.New("the given file type is invalid")
	}

	media := Media{
		file:         file,
		baseFileName: baseFileName,
		ext:          ext,
		name:         filename,
		size:         size,
		path:         filepath.Join(GetUsersImagesDir(), filename),
	}

	return &media, nil
}
