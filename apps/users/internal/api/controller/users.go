package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/qzich/orgserv/pkg/api"
	"github.com/qzich/orgserv/pkg/logger"
)

func NewUser(
	logger logger.Logger,
	reqParser api.RequestParser,
	respSender api.ResponseSender,
) users {
	if logger == nil {
		panic("nil logger")
	}
	if reqParser == nil {
		panic("nil request pareser")
	}
	if respSender == nil {
		panic("nil response sender")
	}
	return users{
		logger:     logger,
		reqParser:  reqParser,
		respSender: respSender,
	}
}

type users struct {
	logger     logger.Logger
	reqParser  api.RequestParser
	respSender api.ResponseSender
}

type UserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u users) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto UserDTO
	if err := u.reqParser.ParseFromBytes(r.Body, &dto); err != nil {
		u.logger.Error(r.Context(), "unable to parse create user request", err.Error())

		u.respSender.SendErrorResponse(w, err)
		return
	}

	u.logger.Debug(r.Context(), dto.Email, dto.Name)

	if len(dto.Name) < 4 || len(dto.Name) > 255 {
		u.respSender.SendErrorResponse(w, fmt.Errorf("name is incorrect: %w", api.ErrValidation))
		return
	}

	if _, err := mail.ParseAddress(dto.Email); err != nil {
		u.respSender.SendErrorResponse(w, fmt.Errorf("email is incorrect: %w", api.ErrValidation))
		return
	}

	w.WriteHeader(http.StatusCreated)
	u.respSender.SendResponse(w, dto)
}
