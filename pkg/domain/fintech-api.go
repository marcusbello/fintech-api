package domain

import "github.com/gin-gonic/gin"

type UserType struct {
	UserName string  `json:"user_name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Account  Account `json:"account omitempty"`
}

type Account struct {
	UserId  string `json:"userId"`
	Balance int    `json:"balance"`
}

type Transaction struct {
	TransferID string `json:"transactionID"`
	From       string `json:"from"`
	To         string `json:"to"`
	Amount     int    `json:"amount"`
}

type FintechUseCase interface {
	LoginUc(c *gin.Context, userName, password string) error
	RegisterUserUc(c *gin.Context, userName, email, password string) (string, error)
	GetUserNameUc(c *gin.Context, userName string) (UserType, error)
	GetAccountUc(c *gin.Context, userName string) (Account, error)
	TransferMoneyUc(c *gin.Context, to, from string, amount int) (Account, error)
}

type FintechRepository interface {
	LoginRepository(c *gin.Context, userName, password string) error
	RegisterUserRepository(c *gin.Context, userName, email, password string) (string, error)
	GetUserNameRepository(c *gin.Context, userName string) (UserType, error)
	GetAccountRepository(c *gin.Context, userName string) (Account, error)
	TransferMoneyRepository(c *gin.Context, to, from string, amount int) (Account, error)
}

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
