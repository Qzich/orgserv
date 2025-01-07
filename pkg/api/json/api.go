package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/qzich/orgserv/pkg/api"
)

type Api struct{}

func (j Api) ParseFromBytes(r io.Reader, v any) error {
	rBytes, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("get client: could not read response body: %w", err)
	}

	if len(rBytes) > 0 {
		if err = json.Unmarshal(rBytes, v); err != nil {
			switch err.(type) {
			case *json.InvalidUnmarshalError:
				return fmt.Errorf("get json: could not unmarshal response body: %w", err)
			}

			return fmt.Errorf("%w: get json: could not unmarshal response body: %w", api.ErrValidation, err)
		}
	}

	return nil
}

type ValidationErrResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (j Api) SendErrorResponse(w http.ResponseWriter, err error) {
	var (
		errorMsg       = "general error"
		httpStatusCode = http.StatusInternalServerError
	)

	if errors.Is(err, api.ErrValidation) {
		errorMsg = "validation error"
		httpStatusCode = http.StatusBadRequest
	}

	if errors.Is(err, api.ErrNotFound) {
		errorMsg = "not found"
		httpStatusCode = http.StatusNotFound
	}

	w.WriteHeader(httpStatusCode)

	data, err := json.Marshal(ValidationErrResponse{Message: errorMsg, Error: err.Error()})
	if err != nil {
		panic(err)
	}

	_, err = w.Write(data)
	if err != nil {
		panic(err)
	}
}

func (j Api) SendResponse(w http.ResponseWriter, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(data)
	if err != nil {
		panic(err)
	}
}
