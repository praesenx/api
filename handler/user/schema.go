package user

import (
    "github.com/gocanto/blog/pkg"
)

type UsersHandler struct {
    Validator  *pkg.Validator
    Repository *Repository
}

type CreatedUser struct {
    UUID string `json:"uuid"`
}

type CreateRequestBag struct {
    FirstName            string `json:"first_name" validate:"required,min=4,max=250"`
    LastName             string `json:"last_name" validate:"required,min=4,max=250"`
    Username             string `json:"username" validate:"required,alphanum,min=4,max=50"`
    DisplayName          string `json:"display_name" validate:"omitempty,min=3,max=255"`
    Email                string `json:"email" validate:"required,email,max=250"`
    Password             string `json:"password" validate:"required,min=8"`
    PublicToken          string `json:"public_token"`
    PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
    Bio                  string `json:"bio" validate:"omitempty"`
    PictureFileName      string `json:"picture_file_name" validate:"omitempty"`
    ProfilePictureURL    string `json:"profile_picture_url" validate:"omitempty,url,max=2048"`
}

type RawCreateRequestBag struct {
    file       []byte
    payload    []byte
    headerName string
}

func (n *RawCreateRequestBag) SetFile(file []byte) {
    n.file = file
}

func (n *RawCreateRequestBag) SetPayload(payload []byte) {
    n.payload = payload
}

func (n *RawCreateRequestBag) SetHeaderName(headerName string) {
    n.headerName = headerName
}

func (n *RawCreateRequestBag) GetFile() []byte {
    return n.file
}

func (n *RawCreateRequestBag) GetPayload() []byte {
    return n.payload
}

func (n *RawCreateRequestBag) GetHeaderName() string {
    return n.headerName
}
