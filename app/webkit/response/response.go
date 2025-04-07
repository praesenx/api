package response

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Response struct {
	Code             int
	Message          string
	Err              error
	ValidationErrors map[string]any
}

func MakeResponse(code int, message string, err error) *Response {
	return &Response{
		Code:             code,
		Message:          message,
		Err:              err,
		ValidationErrors: make(map[string]any),
	}
}

func BadRequest(message string, err error) *Response {
	return MakeResponse(http.StatusBadRequest, message, err)
}

func InternalServerError(message string, err error) *Response {
	return MakeResponse(http.StatusInternalServerError, message, err)
}

func Forbidden(message string, validationErrors map[string]any, err error) *Response {
	return &Response{
		Code:             http.StatusForbidden,
		Message:          message,
		Err:              err,
		ValidationErrors: validationErrors,
	}
}

func Unauthorized(message string, err error) *Response {
	return &Response{
		Code:             http.StatusUnauthorized,
		Message:          message,
		Err:              err,
		ValidationErrors: make(map[string]any),
	}
}

func (e *Response) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}

	return e.Message
}

func (e *Response) Unwrap() error {
	return e.Err
}

func (e *Response) Respond(w http.ResponseWriter) {
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
