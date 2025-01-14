package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qzich/orgserv/pkg/api"
)

//go:generate mockgen --source=../repository/repository.go --destination=./mock_test.go -package=service

func Test_service_CreateUser_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositry := NewMockUsersRepository(ctrl)
	ctx := context.Background()

	service := NewUserService(repositry)
	_, err := service.CreateUser(ctx, "", "", "")
	if !errors.Is(err, api.ErrValidation) {
		t.Error("expected validation error")
	}
}

func Test_service_CreateUser_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositry := NewMockUsersRepository(ctrl)
	ctx := context.Background()

	service := NewUserService(repositry)

	repositry.EXPECT().InsertUser(gomock.Any()).Return(nil)
	_, err := service.CreateUser(ctx, "test", "test@gmail.com", "customer")
	if err != nil {
		t.Errorf("%v", err)
	}
}
