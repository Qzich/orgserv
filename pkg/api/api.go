package api

import (
	"io"
	"net/http"
)

type (
	RequestParser interface {
		ParseFromBytes(io.Reader, any) error
	}

	ResponseSender interface {
		SendErrorResponse(http.ResponseWriter, error)
		SendResponse(http.ResponseWriter, any)
	}
)
