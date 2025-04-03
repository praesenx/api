package media

import (
	"os"
	"path/filepath"
)

func GetStorageDir() string {
	dir, err := os.Getwd()
	folder := Dir

	if err != nil {
		// Handle the error appropriately.
		// Path default to relative path if getting WD fails
		return "./" + folder
	}

	// Resolve the path at runtime.
	return filepath.Join(dir, Dir)
}

func GetImagesDir() string {
	return GetStorageDir() + "/" + ImagesDir
}

func GetUsersImagesDir() string {
	return GetImagesDir() + "/" + UsersDir
}

func GetPostsImagesDir() string {
	return GetImagesDir() + "/" + PostsDir
}
