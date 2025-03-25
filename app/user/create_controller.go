package user

import (
	"encoding/json"
	"github.com/gocanto/blog/app/support"
	"io"
	"log/slog"
	"net/http"
)

func (receiver UsersHandler) create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		slog.Error("Error reading request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var requestBag CreateUsersRequestBag
	if err := json.Unmarshal(body, &requestBag); err != nil {
		slog.Error("Error decoding JSON: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	payload := map[string]interface{}{
		"data": json.RawMessage(body),
	}

	v := support.MakeValidator()

	if _, err := v.Rejects(requestBag); err != nil {
		payload["message"] = err.Error()
		payload["errors"] = v.GetErrors()
	}

	response, err := json.Marshal(payload)
	if err != nil {
		slog.Error("Error generating the response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
