package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oullin/api/env"
	"github.com/oullin/api/pkg"
	"github.com/oullin/api/pkg/media"
	"github.com/oullin/api/pkg/request"
	"github.com/oullin/api/pkg/response"
	"io"
	"mime/multipart"
	"net/http"
)

func (handler UsersHandler) Create(w http.ResponseWriter, r *http.Request) *response.Response {
	var rawRequest RawCreateRequestBag

	multipartRequest, err := request.MakeMultipartRequest(r, &rawRequest)
	defer multipartRequest.Close(nil)

	if err != nil {
		return response.BadRequest("issues creating the request", err)
	}

	err = multipartRequest.ParseRawData(extractData)
	if err != nil {
		return response.BadRequest("NEW: Error getting multipart reader", err)
	}

	var requestBag CreateRequestBag
	if err = json.Unmarshal(rawRequest.payload, &requestBag); err != nil {
		return response.BadRequest("Invalid request payload: malformed JSON", err)
	}

	validate := handler.Validator
	if rejects, err := validate.Rejects(requestBag); rejects {
		return response.Forbidden("Validation failed", validate.GetErrors(), err)
	}

	if result := handler.Repository.FindByUserName(requestBag.Username); result != nil {
		return response.Unprocessable(fmt.Sprintf("user '%s' already exists", requestBag.Username), nil)
	}

	profilePic, err := media.MakeMedia(
		requestBag.Username,
		multipartRequest.GetFile(),
		multipartRequest.GetHeaderName(),
	)

	if err != nil {
		return response.BadRequest("Error handling the given file", err)
	}

	if err := profilePic.Upload(media.GetUsersImagesDir()); err != nil {
		return response.BadRequest("Error saving the given file", err)
	}

	requestBag.PublicToken = r.Header.Get(env.ApiKeyHeader)
	requestBag.PictureFileName = profilePic.GetFileName()
	requestBag.ProfilePictureURL = profilePic.GetFilePath(requestBag.Username)

	created, err := handler.Repository.Create(requestBag)

	if err != nil {
		return response.InternalServerError(err.Error(), err)
	}

	payload := map[string]any{
		"message": "User created successfully!",
		"user": map[string]string{
			"uuid":                created.UUID,
			"picture_file_name":   requestBag.PictureFileName,
			"profile_picture_url": requestBag.ProfilePictureURL,
		},
	}

	return pkg.SendJSON(w, http.StatusCreated, payload)
}

func extractData[T media.MultipartFormInterface](reader *multipart.Reader, data T) error {
	for {
		part, err := reader.NextPart()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch part.FormName() {

		case "data":
			if part.FileName() != "" {
				return errors.New("expected 'data' to be a JSON text field")
			}

			if dataBytes, err := io.ReadAll(part); err != nil {
				return errors.New("Error reading data field" + err.Error())
			} else {
				data.SetPayload(dataBytes)
			}

		case "profile_picture_url":

			if fileBytes, err := io.ReadAll(part); err != nil {
				return errors.New("Error reading file" + err.Error())
			} else {
				data.SetFile(fileBytes)
				data.SetHeaderName(part.FileName())
			}
		}

		if err = part.Close(); err != nil {
			return errors.New("Issue closing the multi-part reader" + err.Error())
		}
	}

	return nil
}
