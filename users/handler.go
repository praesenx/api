package users

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserCreateResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var req UserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Error happened in JSON marshal. Err: %e", err)
	}

	defer r.Body.Close()

	response := UserCreateResponse{
		Message:    "User created successfully",
		StatusCode: http.StatusOK,
		Success:    true,
	}

	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
