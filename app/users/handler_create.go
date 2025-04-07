package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/webkit"
	"github.com/gocanto/blog/app/webkit/media"
	"github.com/gocanto/blog/app/webkit/request"
	"github.com/gocanto/blog/app/webkit/response"
	"io"
	"mime/multipart"
	"net/http"
)

func (handler UserHandler) Create(w http.ResponseWriter, r *http.Request) *response.Response {
	var profilePhoto RawCreateRequestBag

	req, err := request.MakeMultipartRequest(r, &profilePhoto)
	defer req.Close(nil)

	if err != nil {
		fmt.Println("1 ---> ", err)
		return response.BadRequest("issues creating the request", err)
	}

	err = req.ParseRawData(extractData)
	if err != nil {
		fmt.Println("---> ", err)
		return response.BadRequest("NEW: Error getting multipart reader", err)
	}

	profilePic, err := media.MakeMedia(req.GetFile(), req.GetHeaderName())

	if err != nil {
		return response.BadRequest("Error handling the given file", err)
	}

	if err := profilePic.Write(); err != nil {
		return response.BadRequest("Error saving the given file", err)
	}

	var requestBag CreateRequestBag
	if err = json.Unmarshal(profilePhoto.payload, &requestBag); err != nil {
		return response.BadRequest("Invalid request payload: malformed JSON", err)
	}

	validate := handler.Validator
	if rejects, err := validate.Rejects(requestBag); rejects {
		return response.Forbidden("Validation failed", validate.GetErrors(), err)
	}

	if result := handler.Repository.FindByUserName(requestBag.Username); result != nil {
		return response.Forbidden(
			fmt.Sprintf("user '%s' already exists", requestBag.Username),
			map[string]any{},
			nil,
		)
	}

	requestBag.PublicToken = r.Header.Get(env.ApiKeyHeader)
	created, err := handler.Repository.Create(requestBag)

	if err != nil {
		return response.InternalServerError(err.Error(), err)
	}

	payload := map[string]any{
		"message": "User created successfully!",
		"user":    map[string]string{"uuid": created.UUID},
		//"data":    json.RawMessage(body),
	}

	return webkit.SendJSON(w, http.StatusCreated, payload)
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
