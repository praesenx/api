package users

import (
	"encoding/json"
	"github.com/gocanto/blog/support"
	"log"
	"log/slog"
	"net/http"
)

type UserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type UserCreateResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type UserHandler struct {
	Validator support.Validator
}

func Create(w http.ResponseWriter, r *http.Request) {
	var userRequest UserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		// --- Todo: Return bad request: 400
		slog.Error("Error happened in JSON marshal. Err: %e", err)
	}

	defer r.Body.Close()

	response := make(map[string]interface{})

	v := support.MakeValidator()

	if _, err := v.Rejects(userRequest); err != nil {
		response["errors"] = v.GetErrors()
	}

	jsonResp, err := json.Marshal(response)
	if err != nil {
		// --- Todo: Return server error: 500
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
