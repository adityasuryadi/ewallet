package services

import (
	"context"

	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/adityasuryadi/ewallet/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	RegisterRepository interfaces.IRegisterRepository
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	password := request.Password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashPassword)
	resp := request

	err = s.RegisterRepository.InsertUser(&resp)
	if err != nil {
		return nil, err
	}

	resp.Password = ""
	return request, nil
}
