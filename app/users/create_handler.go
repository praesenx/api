package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/kernel"
	"github.com/gocanto/blog/app/media"
	"io"
	"log/slog"
	"net/http"
)

//const maxFileSize = 10 * 1024 * 1024 // 10 MB
//var allowedExtensions = []string{".jpg", ".jpeg", ".png"}

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

func (handler HandleUsers) Create(w http.ResponseWriter, r *http.Request) *kernel.HttpError {
	body, err := io.ReadAll(r.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("Issue closing the request body", err)
		}
	}(r.Body)

	if err != nil {
		return kernel.BadRequest("Invalid request payload: cannot read body", err)
	}

	fmt.Println("Raw body length:", len(body))
	// Reset the request body so it can be read again.
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// Get the multipart reader.
	mr, err := r.MultipartReader()
	if err != nil {
		return kernel.BadRequest("Error getting multipart reader", err)
	}

	var fileBytes []byte
	var dataBytes []byte
	var fileHeaderName string

	for {
		part, err := mr.NextPart()

		if err == io.EOF {
			break
		}

		if err != nil {
			return kernel.BadRequest("Error reading multipart parts", err)
		}

		// Check which part we got.
		switch part.FormName() {

		case "data":
			// Ensure this is a text field (not a file).
			if part.FileName() != "" {
				return kernel.BadRequest("Expected 'data' to be a JSON text field", err)
			}

			dataBytes, err = io.ReadAll(part)
			if err != nil {
				return kernel.BadRequest("Error reading data field", err)
			}

			fmt.Println("Received data field:", string(dataBytes))

		case "profile_picture_url":

			fileBytes, err = io.ReadAll(part)
			if err != nil {
				return kernel.BadRequest("Error reading file", err)
			}
			fileHeaderName = part.FileName()
			fmt.Printf("Received file part: %d bytes\n", len(fileBytes))
		default:
			fmt.Println("Ignoring unexpected part:", part.FormName())
		}

		if err = part.Close(); err != nil {
			slog.Error("Issue closing the multi-part reader", err)
		}
	}

	// --- Save the file using fileBytes ---
	profilePic, err := media.MakeMedia(fileBytes, fileHeaderName)

	if err != nil {
		return kernel.BadRequest("Error handling the given file", err)
	}

	if err := profilePic.Write(); err != nil {
		return kernel.BadRequest("Error saving the given file", err)
	}

	var requestBag CreateRequestBag
	if err = json.Unmarshal(dataBytes, &requestBag); err != nil {
		return kernel.BadRequest("Invalid request payload: malformed JSON", err)
	}

	validate := handler.Validator
	if rejects, err := validate.Rejects(requestBag); rejects {
		return kernel.RespondWithErrors("Validation failed", validate.GetErrors(), err)
	}

	if result := handler.Repository.FindByUserName(requestBag.Username); result != nil {
		return kernel.RespondWithErrors(
			fmt.Sprintf("user '%s' already exists", requestBag.Username),
			map[string]any{},
			nil,
		)
	}

	requestBag.PublicToken = r.Header.Get(env.ApiKeyHeader)
	created, err := handler.Repository.Create(requestBag)

	if err != nil {
		return kernel.InternalServerError(err.Error(), err)
	}

	payload := map[string]any{
		"message": "User created successfully!",
		"user":    map[string]string{"uuid": created.UUID},
		//"data":    json.RawMessage(body),
	}

	return kernel.SendJSON(w, http.StatusCreated, payload)
}
