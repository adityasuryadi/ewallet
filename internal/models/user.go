package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int       `json:"id" gorm:"column:id;autoIncrement;primary_key;"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(255);" validate:"required"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(255);" validate:"required,email"`
	Password    string    `json:"password,omitempty" gorm:"column:password;type:varchar(255);" validate:"required,min=6"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(255);" validate:"required"`
	FullName    string    `json:"full_name" gorm:"column:full_name;type:varchar(255);" validate:"required"`
	Address     string    `json:"address" gorm:"column:address;type:varchar(255);" validate:"required"`
	Dob         string    `json:"dob" gorm:"column:dob;type:date;" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (u *User) TableName() string {
	return "users"
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  int       `json:"id"`
	UserID              int       `json:"user_id"`
	Token               string    `json:"token"`
	TokenExpired        time.Time `json:"token_expired"`
	RefreshToken        string    `json:"refresh_token"`
	RefreshTokenExpired time.Time `json:"refresh_token_expired"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (u *UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
