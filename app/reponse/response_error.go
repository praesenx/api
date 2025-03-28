package reponse

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type ResponseError struct {
	Code             int
	Message          string
	Err              error
	ValidationErrors map[string]any
}

func MakeResponseError(code int, message string, err error) *ResponseError {
	return &ResponseError{
		Code:             code,
		Message:          message,
		Err:              err,
		ValidationErrors: make(map[string]any),
	}
}

func MakeBadRequest(message string, err error) *ResponseError {
	return MakeResponseError(http.StatusBadRequest, message, err)
}

func MakeInternalServerError(message string, err error) *ResponseError {
	return MakeResponseError(http.StatusInternalServerError, message, err)
}

func MakeValidationError(message string, validationErrors map[string]any, err error) *ResponseError {
	return &ResponseError{
		Code:             http.StatusBadRequest,
		Message:          message,
		Err:              err,
		ValidationErrors: validationErrors,
	}
}

func (e *ResponseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}

	return e.Message
}

func (e *ResponseError) Unwrap() error {
	return e.Err
}

func (e *ResponseError) Respond(w http.ResponseWriter) {
	slog.Error("HTTP Error", "status", e.Code, "message", e.Message, "error", e.Err, "validation_errors", e.ValidationErrors)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff") // Basic security header
	w.WriteHeader(e.Code)

	payload := map[string]any{
		"message": e.Message,
	}

	if len(e.ValidationErrors) > 0 {
		payload["errors"] = e.ValidationErrors
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("Error encoding error response", "encode_error", err, "original_error", e)
		_, _ = fmt.Fprintf(w, `{"message":"Error generating error response"}`)
	}
}
