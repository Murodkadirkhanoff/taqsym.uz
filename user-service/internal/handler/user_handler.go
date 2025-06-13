package handler

import (
	"context"
	"errors"

	authpb "github.com/Murodkadirkhanoff/taqsym.uz/proto/auth"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/domain"
)

type UserHandler struct {
	authpb.UnimplementedAuthServiceServer
	uc domain.UserUseCase
}

func NewUserHandler(uc domain.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	token, err := h.uc.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &authpb.LoginResponse{Token: token}, nil
}

func (h *UserHandler) Register(c context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	user := domain.User{
		Email:    req.GetEmail(),
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}

	if err := h.uc.Register(c, &user); err != nil {
		return nil, errors.New(err.Error())
	}
	return &authpb.RegisterResponse{
		Id:      user.ID,
		Message: "User Created Successfully",
	}, nil
}

func (h *UserHandler) Profile(c context.Context, req *authpb.ProfileRequest) (*authpb.ProfileResponse, error) {
	user, err := h.uc.Profile(c, int(req.GetId()))

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &authpb.ProfileResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil

}
