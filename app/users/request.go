package users

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type Request struct {
	raw        *http.Request
	body       io.ReadCloser
	parsedBody []byte
	data       []byte
	payload    map[string]interface{}
}

func MakeUsersRequest(request *http.Request, bag *CreateRequestBag) (*Request, error) {
	body := request.Body
	parsedBody, err := io.ReadAll(body)
	//instance := Request{}

	if err != nil {
		slog.Error("Error reading the request body JSON: %v", err)
		return nil, err
	}

	//instance.raw = request
	//instance.body = body
	//instance.parsedBody = parsedBody
	//var requestBag CreateRequestBag

	if err = json.Unmarshal(parsedBody, bag); err != nil {
		slog.Error("Error decoding the request body JSON: %v", err)
		return nil, err
	}

	data := json.RawMessage(parsedBody)

	return &Request{
		raw:        request,
		body:       body,
		parsedBody: parsedBody,
		data:       data,
	}, nil
}

//func (current Request) validate(validate func(validator support.Validator) (bool, error)) Request {
//	validate(validate({
//
//	}))
//}

func (current Request) Close() {
	err := current.body.Close()

	if err != nil {
		slog.Error("Error closing the request body: " + err.Error())
	}
}
