package auth

import (
	"go-clean/middlewares"
	"go-clean/modules/user"
	"time"

	"gorm.io/gorm"
)

type TokenType string

const (
	Access        TokenType = "ACCESS"
	Refresh       TokenType = "REFRESH"
	ResetPassword TokenType = "RESET_PASSWORD"
	VerifyEmail   TokenType = "VERIFY_EMAIL"
)

type ModelToken struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	Token       string         `gorm:"unique;not null"`
	Type        TokenType      `gorm:"type:token_type;not null"`
	Expires     time.Time      `gorm:"not null"`
	Blacklisted bool           `gorm:"-" json:"-"`
	UserID      uint           `gorm:"not null"`
	User        user.ModelUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (ModelToken) TableName() string {
	return "tokens"
}

var ValidateLogin = middlewares.Validator{
	Body: &RequestLogin{},
}
var ValidateRefreshToken = middlewares.Validator{
	Body: &RequestRefreshToken{},
}

var ValidateLoginEmail = middlewares.Validator{
	Body: &RequestLoginEmail{},
}

type RequestLoginEmail struct {
	Credential string `json:"credential" validate:"required,min=3,max=32"`
	Password   string `json:"password" validate:"required,min=8"`
}

type RequestLogin struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8"`
}

type RequestRefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=8"`
}

type ResponseToken struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}

type ResponseAuthToken struct {
	Access  ResponseToken `json:"access"`
	Refresh ResponseToken `json:"refresh"`
}
