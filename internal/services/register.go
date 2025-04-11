package services

import (
	"context"

	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/adityasuryadi/ewallet/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	log = helpers.Logger
)

type RegisterService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	password := request.Password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error hash password ", err)
		return nil, err
	}
	request.Password = string(hashPassword)
	resp := request

	err = s.UserRepository.InsertUser(&resp)
	if err != nil {
		log.Error("error insert user ", err)
		return nil, err
	}

	resp.Password = ""
	return request, nil
}
