package users

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type CreateRequestBag struct {
	FirstName            string `json:"first_name" validate:"required,min=4,max=250"`
	LastName             string `json:"last_name" validate:"required,min=4,max=250"`
	Username             string `json:"username" validate:"required,alphanum,min=4,max=50"`
	DisplayName          string `json:"display_name" validate:"omitempty,min=4,max=255"`
	Email                string `json:"email" validate:"required,email,max=250"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	Bio                  string `json:"bio" validate:"omitempty"`
	ProfilePictureURL    string `json:"profile_picture_url" validate:"omitempty,url,max=2048"`
	RawRequest           []byte
}

func (handler Handler) create(w http.ResponseWriter, r *http.Request) {
	//request, err := MakeUsersRequest(r)
	//
	//if err != nil {
	//	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	//	return
	//}
	//
	//request.Close()

	rawBody := r.Body
	body, err := io.ReadAll(rawBody)
	defer rawBody.Close()

	if err != nil {
		slog.Error("Error reading request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var requestBag CreateRequestBag
	if err = json.Unmarshal(body, &requestBag); err != nil {
		slog.Error("Error decoding JSON: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	payload := map[string]interface{}{
		"data": json.RawMessage(body),
	}

	v := handler.Validator

	if _, err = v.Rejects(requestBag); err != nil {
		payload["message"] = err.Error()
		payload["errors"] = v.GetErrors()
	}

	// --- repo call here
	//createdUser, err := handler.Repository.Create(requestBag)

	response, err := json.Marshal(payload)
	if err != nil {
		slog.Error("Error generating the response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
