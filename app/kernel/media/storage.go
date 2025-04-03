package media

import "mime/multipart"

const UsersDir = "users"
const PostsDir = "posts"
const ImagesDir = "images"
const Dir = "storage"

type Storage interface {
	Save(file multipart.File, filename string) (string, error)
}
