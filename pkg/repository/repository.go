package repository

import (
	"fintech-api/pkg/domain"
	"github.com/gin-gonic/gin"
)

type FintechRepository interface {
	LoginRepository(c *gin.Context, userName, password string) error
	RegisterUserRepository(c *gin.Context, userName, email, password string) (string, error)
	GetUserRepository(c *gin.Context, userName string) (domain.UserType, error)
	GetAccountRepository(c *gin.Context, userName string) (domain.Account, error)
	AddMoneyRepository(c *gin.Context, to string, amount int) (domain.Account, error)
	RemoveMoneyRepository(c *gin.Context, from string, amount int) (domain.Account, error)
	AddToTransaction(c *gin.Context, from, to string, amount int) error
}
