package users

import (
	"encoding/json"
	"fmt"
	"github.com/gocanto/blog/app/reponse"
	"github.com/gocanto/blog/app/support"
	"io"
	"net/http"
)

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
	ProfilePictureURL    string `json:"profile_picture_url" validate:"omitempty,url,max=2048"`
}

func (handler HandleUsers) Create(w http.ResponseWriter, r *http.Request) *reponse.ResponseError {
	body, err := io.ReadAll(r.Body)
	defer support.CloseRequestBody(r)

	if err != nil {
		return reponse.MakeBadRequest("Invalid request payload: cannot read body", err)
	}

	var requestBag CreateRequestBag
	if err = json.Unmarshal(body, &requestBag); err != nil {
		return reponse.MakeBadRequest("Invalid request payload: malformed JSON", err)
	}

	validate := handler.Validator
	if rejects, err := validate.Rejects(requestBag); rejects {
		return reponse.MakeValidationError("Validation failed", validate.GetErrors(), err)
	}

	if result := handler.Repository.FindByUserName(requestBag.Username); result != nil {
		return reponse.MakeValidationError(
			fmt.Sprintf("user '%s' already exists", requestBag.Username),
			map[string]any{},
			nil,
		)
	}

	requestBag.PublicToken = r.Header.Get(support.ApiKeyHeader)
	created, err := handler.Repository.Create(requestBag)

	if err != nil {
		return reponse.MakeInternalServerError(err.Error(), err)
	}

	payload := map[string]any{
		"message": "User created successfully!",
		"user":    map[string]string{"uuid": created.UUID},
		//"data":    json.RawMessage(body),
	}

	return reponse.SendJSON(w, http.StatusCreated, payload)
}
