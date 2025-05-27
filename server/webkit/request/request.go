package request

import (
	"errors"
	"fmt"
	"github.com/gocanto/blog/server/webkit/media"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
)

type Request struct {
	baseRequest      *http.Request
	isMultipart      bool
	multipartReader  *multipart.Reader
	multiPartRawData media.MultipartFormInterface
}

func MakeMultipartRequest[T media.MultipartFormInterface](r *http.Request, rawData T) (*Request, error) {
	reader, err := r.MultipartReader()

	if err != nil {
		return nil, errors.New("the isMultipart form reader is invalid")
	}

	return &Request{
		baseRequest:      r,
		isMultipart:      true,
		multiPartRawData: rawData,
		multipartReader:  reader,
	}, nil
}

func (req *Request) Close(message *string) {
	m := "Issue closing the request body"

	if message == nil {
		message = &m
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error(m, err)
		}
	}(req.baseRequest.Body)
}

func (req *Request) ParseRawData(callback func(reader *multipart.Reader, data media.MultipartFormInterface) error) error {
	fmt.Println(fmt.Sprintf("dd: %+v", req))
	if req.multipartReader == nil {
		return errors.New("1) invalid isMultipart form")
	}

	if req.multiPartRawData == nil {
		return errors.New("2) invalid isMultipart form request")
	}

	result := callback(req.multipartReader, req.multiPartRawData)

	if result != nil {
		return errors.New("3) invalid isMultipart form parsing: " + result.Error())
	}

	return nil
}

func (req *Request) GetFile() []byte {
	return req.multiPartRawData.GetFile()
}

func (req *Request) GetHeaderName() string {
	return req.multiPartRawData.GetHeaderName()
}
