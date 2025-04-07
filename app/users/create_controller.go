package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocanto/blog/app/controller"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/media"
	"io"
	"log/slog"
	"mime/multipart"
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

func (handler UserController) Create(w http.ResponseWriter, r *http.Request) *controller.HttpError {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("Issue closing the request body", err)
		}
	}(r.Body)

	// Get the multipart reader.
	mr, err := r.MultipartReader()
	if err != nil {
		return controller.BadRequest("Error getting multipart reader", err)
	}

	var profilePhoto UserProfilePhoto

	if err := extractData(mr, &profilePhoto); err != nil {
		return controller.BadRequest("Error extracting data", err)
	}

	// --- Save the file using fileBytes ---
	profilePic, err := media.MakeMedia(profilePhoto.file, profilePhoto.headerName)

	fmt.Println("--->", len(profilePhoto.file), "---> err: ", err)
	if err != nil {
		return controller.BadRequest("Error handling the given file", err)
	}

	if err := profilePic.Write(); err != nil {
		return controller.BadRequest("Error saving the given file", err)
	}

	var requestBag CreateRequestBag
	if err = json.Unmarshal(profilePhoto.payload, &requestBag); err != nil {
		return controller.BadRequest("Invalid request payload: malformed JSON", err)
	}

	validate := handler.Validator
	if rejects, err := validate.Rejects(requestBag); rejects {
		return controller.RespondWithErrors("Validation failed", validate.GetErrors(), err)
	}

	if result := handler.Repository.FindByUserName(requestBag.Username); result != nil {
		return controller.RespondWithErrors(
			fmt.Sprintf("user '%s' already exists", requestBag.Username),
			map[string]any{},
			nil,
		)
	}

	requestBag.PublicToken = r.Header.Get(env.ApiKeyHeader)
	created, err := handler.Repository.Create(requestBag)

	if err != nil {
		return controller.InternalServerError(err.Error(), err)
	}

	payload := map[string]any{
		"message": "User created successfully!",
		"user":    map[string]string{"uuid": created.UUID},
		//"data":    json.RawMessage(body),
	}

	return controller.SendJSON(w, http.StatusCreated, payload)
}

func extractData(reader *multipart.Reader, data *UserProfilePhoto) error {
	var fileBytes []byte
	var dataBytes []byte
	var fileHeaderName string

	for {
		part, err := reader.NextPart()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		// Check which part we got.
		switch part.FormName() {

		case "data":
			if part.FileName() != "" {
				return errors.New("expected 'data' to be a JSON text field")
			}

			dataBytes, err = io.ReadAll(part)
			if err != nil {
				return errors.New("Error reading data field" + err.Error())
			}

			data.payload = dataBytes
			fmt.Println("Received data field:", string(dataBytes))

		case "profile_picture_url":

			fileBytes, err = io.ReadAll(part)
			if err != nil {
				return controller.BadRequest("Error reading file", err)
			}

			fileHeaderName = part.FileName()
			fmt.Printf("Received file name: %s\n", fileHeaderName)
			fmt.Printf("Received file part: %d bytes\n", len(fileBytes))

			data.file = fileBytes
			data.headerName = fileHeaderName
		default:
			fmt.Println("Ignoring unexpected part:", part.FormName())
		}

		if err = part.Close(); err != nil {
			return errors.New("Issue closing the multi-part reader" + err.Error())
		}
	}

	return nil
}
