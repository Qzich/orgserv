package controller

import (
	"fmt"
	"net/http"

	"github.com/qzich/orgserv/pkg/api"
	"github.com/qzich/orgserv/pkg/api/dto"
	"github.com/qzich/orgserv/pkg/logger"
	"github.com/qzich/orgserv/pkg/service"
	"github.com/qzich/orgserv/pkg/uuid"
)

type (
	users struct {
		logger     logger.Logger
		reqParser  api.RequestParser
		respSender api.ResponseSender
		srv        service.UsersService
	}
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

func (u users) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserDTO
	if err := u.reqParser.ParseFromBytes(r.Body, &dto); err != nil {
		u.logger.Error(r.Context(), "unable to parse create user request", err.Error())

		u.respSender.SendErrorResponse(w, err)
		return
	}

	u.logger.Debug(r.Context(), dto.Email, dto.Name, dto.Kind)

	// fmt.Printf("%#v", dto)

	user, err := u.srv.CreateUser(r.Context(), dto.Name, dto.Email, dto.Kind, dto.Password)
	if err != nil {
		u.respSender.SendErrorResponse(w, err)
		return
	}

	dto.ID = user.ID().String()
	dto.CreatedAt = user.CreatedAt()
	dto.UpdatedAt = user.UpdatedAt()

	w.WriteHeader(http.StatusCreated)
	u.respSender.SendResponse(w, dto)
}

func (u users) GetUser(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserDTO

	id := r.PathValue("id")

	userId, err := uuid.FromString(id)
	if err != nil {
		u.respSender.SendErrorResponse(w, fmt.Errorf("id is incorrect: %w", api.ErrValidation))
		return
	}

	user, err := u.srv.GetUser(r.Context(), userId)
	if err != nil {
		u.respSender.SendErrorResponse(w, err)
		return
	}

	dto.ID = user.ID().String()
	dto.Name = user.Name()
	dto.Email = user.Email()
	dto.Kind = user.Kind().String()
	dto.CreatedAt = user.CreatedAt()
	dto.UpdatedAt = user.UpdatedAt()

	u.respSender.SendResponse(w, dto)
}

func (u users) UsersList(w http.ResponseWriter, r *http.Request) {
	list := make([]dto.UserDTO, 0)

	users, err := u.srv.AllUsers(r.Context())
	if err != nil {
		u.respSender.SendErrorResponse(w, err)
		return
	}

	for _, user := range users {
		var dto dto.UserDTO

		dto.ID = user.ID().String()
		dto.Name = user.Name()
		dto.Email = user.Email()
		dto.Kind = user.Kind().String()
		dto.CreatedAt = user.CreatedAt()
		dto.UpdatedAt = user.UpdatedAt()

		list = append(list, dto)

	}

	u.respSender.SendResponse(w, list)
}
