package storage

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
		return "./" + folder //
	}

	// Resolve the path at runtime.
	return filepath.Join(dir, Dir)
}

func GetImageDir() string {
	return GetStorageDir() + "/" + ImagesDir
}

func GetUsersDir() string {
	return GetStorageDir() + "/" + UsersDir
}

func GetPostsDir() string {
	return GetStorageDir() + "/" + PostsDir
}
