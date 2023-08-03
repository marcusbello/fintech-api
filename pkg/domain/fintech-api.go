package domain

import "github.com/gin-gonic/gin"

type User struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}

// UserType example
type UserType struct {
	UserName string  `json:"user_name"`
	Email    string  `json:"email"`
	Account  Account `json:"account,omitempty"`
} //@name UserProfile

// Account example
type Account struct {
	UserName string `json:"user_name,omitempty"`
	Balance  int    `json:"balance"`
} //@name Account

type AccountType struct {
	Balance int `json:"balance"`
}

type Transaction struct {
	TransferID string `json:"transactionID"`
	From       string `json:"from"`
	To         string `json:"to"`
	Amount     int    `json:"amount"`
}

// PingPong example
type PingPong struct {
	Data string `json:"data"`
} //@name PingPong

type FintechUseCase interface {
	LoginUc(c *gin.Context, userName, password string) error
	RegisterUserUc(c *gin.Context, userName, email, password string) (string, error)
	GetUserUc(c *gin.Context, userName string) (UserType, error)
	GetAccountUc(c *gin.Context, userName string) (Account, error)
	TransferMoneyUc(c *gin.Context, from, to string, amount int) (Account, error)
	AddMoneyUc(c *gin.Context, to string, amount int) (Account, error)
	RemoveMoneyUc(c *gin.Context, from string, amount int) (Account, error)
}

// LoginRequest example
type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
} //@name LoginRequest

// RegisterRequest example
type RegisterRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required"`
} //@name RegisterRequest

// TransferRequest example
type TransferRequest struct {
	From   string `json:"from" swaggerignore:"true"`
	To     string `json:"to" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
} //@name TransferRequest
