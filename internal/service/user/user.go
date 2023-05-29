package user

import (
	"context"

	"github.com/i-akbarshoh/auth-service/internal/pkg/proto"
	"github.com/i-akbarshoh/auth-service/internal/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repository
	proto.AuthServer
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) SignUp(ctx context.Context, user *proto.SignUpModel) ( *proto.Empty, error) {
	password, err := utils.GeneratePasswordHash(user.Password)
	if err != nil {
		return &proto.Empty{}, err
	}
	create := Create{
		ID:       user.Id,
		Name:     user.Name,
		Password: string(password),
		Role:     user.Role,
		Age:      user.Age,
		Email:    user.Email,
	}

	if err := s.repo.Create(ctx, create); err != nil {
		return &proto.Empty{}, err
	}

	return &proto.Empty{}, nil
}

func (s *service) Login(ctx context.Context, get *proto.LoginModel) (*proto.Empty, error) {
	res, err := s.repo.Get(ctx, get.Email)
	if err != nil {
		return &proto.Empty{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(get.Password)); err != nil {
		return &proto.Empty{}, err
	}

	return &proto.Empty{}, nil
}
