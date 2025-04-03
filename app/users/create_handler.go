package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocanto/blog/app/storage"
	"github.com/gocanto/blog/app/support"
	"github.com/google/uuid"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	//"bytes"
	//"encoding/json"
	//"fmt"
	"github.com/gocanto/blog/app/kernel"
	"io"

	//"github.com/gocanto/blog/app/support"
	//"io"
	"net/http"
)

var storageDir string

const maxFileSize = 10 * 1024 * 1024 // 10 MB
var allowedExtensions = []string{".jpg", ".jpeg", ".png"}

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

func (handler HandleUsers) Create(w http.ResponseWriter, r *http.Request) *kernel.HttpException {
	body, err := io.ReadAll(r.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("Issue closing the request body", err)
		}
	}(r.Body)

	if err != nil {
		return kernel.MakeBadRequestException("Invalid request payload: cannot read body", err)
	}

	fmt.Println("Raw body length:", len(body))
	// Reset the request body so it can be read again.
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// Get the multipart reader.
	mr, err := r.MultipartReader()
	if err != nil {
		return kernel.MakeBadRequestException("Error getting multipart reader", err)
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
			return kernel.MakeBadRequestException("Error reading multipart parts", err)
		}

		// Check which part we got.
		switch part.FormName() {

		case "data":
			// Ensure this is a text field (not a file).
			if part.FileName() != "" {
				return kernel.MakeBadRequestException("Expected 'data' to be a JSON text field", err)
			}

			dataBytes, err = io.ReadAll(part)
			if err != nil {
				return kernel.MakeBadRequestException("Error reading data field", err)
			}

			fmt.Println("Received data field:", string(dataBytes))

		case "profile_picture_url":

			fileBytes, err = io.ReadAll(part)
			if err != nil {
				return kernel.MakeBadRequestException("Error reading file", err)
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
	if len(fileBytes) > 0 {
		ext := strings.ToLower(filepath.Ext(fileHeaderName))
		filename := uuid.New().String() + ext
		filePath := filepath.Join(storage.GetUsersImagesDir(), filename)

		err = os.WriteFile(filePath, fileBytes, 0644)
		fmt.Println("---> File Path:", filePath) // Add this line
		//dst, err := os.Create(filePath)

		if err != nil {
			fmt.Println("Error writing file:", filename, " <---", err)
			return kernel.MakeInternalServerException("Error saving the file", err)
		}
		fmt.Println("---> Storage folder:", storage.GetUsersImagesDir())
		fmt.Println("---> File Path:", filePath)
	}
	// --- End of file saving ---

	//var requestBag CreateRequestBag
	//if err = json.Unmarshal(dataBytes, &requestBag); err != nil {
	//	return kernel.MakeBadRequestException("Invalid request payload: malformed JSON", err)
	//}

	//err = r.ParseMultipartForm(maxFileSize)
	//if err != nil {
	//	fmt.Println("--> ", err, " <---")
	//	return kernel.MakeBadRequestException("Error parsing multipart form:", err)
	//}

	// Correctly assign the three return values
	//file, fileHeader, err := r.FormFile("profile_picture_url")
	//if err != nil {
	//	return kernel.MakeBadRequestException("Error retrieving the file", err)
	//}
	//defer func(file multipart.File) {
	//	err := file.Close()
	//	if err != nil {
	//		slog.Error("Issue closing the multipart file", err)
	//	}
	//}(file)
	//// You can now access information from fileHeader, e.g., fileHeader.Filename
	//fmt.Println("Uploaded filename:", fileHeader.Filename)
	//
	//ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	//isValidExtension := false
	//for _, allowedExt := range allowedExtensions {
	//	if ext == allowedExt {
	//		isValidExtension = true
	//		break
	//	}
	//}
	//
	//if !isValidExtension {
	//	return kernel.MakeBadRequestException("Invalid file extension", err)
	//}

	//err = os.MkdirAll(storageDir, os.ModePerm)
	//if err != nil {
	//	//http.Error(w, fmt.Sprintf("Error creating storage directory: %v", err), http.StatusInternalServerError)
	//	return kernel.MakeInternalServerException("Error creating storage directory", err)
	//}

	//filename := uuid.New().String() + ext
	//filePath := filepath.Join(storage.GetUsersImagesDir(), filename)
	//
	//dst, err := os.Create(filePath)
	//if err != nil {
	//	return kernel.MakeInternalServerException("Error creating destination file", err)
	//}
	//defer func(dst *os.File) {
	//	err := dst.Close()
	//	if err != nil {
	//		slog.Error("Issue closing the destination file", err)
	//	}
	//}(dst)
	//
	//_, err = io.Copy(dst, file)
	//if err != nil {
	//	return kernel.MakeInternalServerException("Error saving the file", err)
	//}

	// ------
	//fmt.Println("---> Storage folder:", storage.GetUsersImagesDir())
	//fmt.Println("---> File Path:", filePath)
	// ------

	var requestBag CreateRequestBag
	if err = json.Unmarshal(dataBytes, &requestBag); err != nil {
		return kernel.MakeBadRequestException("Invalid request payload: malformed JSON", err)
	}

	validate := handler.Validator
	if rejects, err := validate.Rejects(requestBag); rejects {
		return kernel.MakeValidationException("Validation failed", validate.GetErrors(), err)
	}

	if result := handler.Repository.FindByUserName(requestBag.Username); result != nil {
		return kernel.MakeValidationException(
			fmt.Sprintf("user '%s' already exists", requestBag.Username),
			map[string]any{},
			nil,
		)
	}

	requestBag.PublicToken = r.Header.Get(support.ApiKeyHeader)
	created, err := handler.Repository.Create(requestBag)

	if err != nil {
		return kernel.MakeInternalServerException(err.Error(), err)
	}

	payload := map[string]any{
		"message": "User created successfully!",
		"user":    map[string]string{"uuid": created.UUID},
		//"data":    json.RawMessage(body),
	}

	return kernel.SendJSON(w, http.StatusCreated, payload)
}
