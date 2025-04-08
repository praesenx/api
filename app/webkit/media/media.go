package media

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const Dir = "media"
const UsersDir = "users"
const PostsDir = "posts"
const StorageDir = "storage"

var maxFileSize = int64(50 * 1024 * 1024) // 50 MB in bytes
var allowedExtensions = []string{".jpg", ".jpeg", ".png"}

type Media struct {
	file         []byte
	baseFileName string
	ext          string
	name         string
	path         string
	size         int64
}

func (m *Media) Upload() error {
	err := os.WriteFile(m.path, m.file, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (m *Media) GetFileName() string {
	return m.name
}
func (m *Media) GetFilePath(prefix string) string {
	filePath := m.path

	dir := filepath.Dir(filePath)
	ext := filepath.Ext(filePath)
	base := filepath.Base(filePath)
	name := strings.TrimSuffix(base, ext)
	newName := fmt.Sprintf("%s-%s%s", prefix, name, ext)

	return filepath.Join(dir, newName)
}

func (m *Media) GetHeaderName() string {
	return m.baseFileName
}

func (m *Media) GetExtension() string {
	return m.ext
}
