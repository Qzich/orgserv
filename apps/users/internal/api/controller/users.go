package controller

import (
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/qzich/orgserv/pkg/api"
	"github.com/qzich/orgserv/pkg/logger"
	"github.com/qzich/orgserv/pkg/service"
)

func NewUser(
	logger logger.Logger,
	reqParser api.RequestParser,
	respSender api.ResponseSender,
	srv service.UsersService,
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
	if srv == nil {
		panic("nil user service")
	}
	return users{
		logger:     logger,
		reqParser:  reqParser,
		respSender: respSender,
		srv:        srv,
	}
}

type users struct {
	logger     logger.Logger
	reqParser  api.RequestParser
	respSender api.ResponseSender
	srv        service.UsersService
}

type UserDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Kind      string    `json:"kind"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

	if dto.Kind != "support" && dto.Kind != "customer" {
		u.respSender.SendErrorResponse(w, fmt.Errorf("kind is incorrect: %w", api.ErrValidation))
		return
	}

	// fmt.Printf("%#v", dto)

	user, err := u.srv.CreateUser(r.Context(), dto.Name, dto.Email, dto.Kind)
	if err != nil {
		u.respSender.SendErrorResponse(w, err)
		return
	}

	dto.ID = user.ID.String()
	dto.Kind = user.Kind
	dto.CreatedAt = user.CreatedAt
	dto.UpdatedAt = user.UpdatedAt

	w.WriteHeader(http.StatusCreated)
	u.respSender.SendResponse(w, dto)
}
