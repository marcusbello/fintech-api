package usecase

import (
	"fintech-api/pkg/domain"
	"github.com/gin-gonic/gin"
)

type fintechUc struct {
	fintechRepo domain.FintechRepository
}

func (u fintechUc) LoginUc(c *gin.Context, userName, password string) error {
	return u.fintechRepo.LoginRepository(c, userName, password)
}

func (u fintechUc) RegisterUserUc(c *gin.Context, userName, email, password string) (string, error) {
	return u.fintechRepo.RegisterUserRepository(c, userName, email, password)
}

func (u fintechUc) GetUserNameUc(c *gin.Context, userName string) (domain.UserType, error) {
	return u.fintechRepo.GetUserNameRepository(c, userName)
}

func (u fintechUc) GetAccountUc(c *gin.Context, userName string) (domain.Account, error) {
	return u.fintechRepo.GetAccountRepository(c, userName)
}

func (u fintechUc) TransferMoneyUc(c *gin.Context, to, from string, amount int) (domain.Account, error) {
	return u.fintechRepo.TransferMoneyRepository(c, to, from, amount)
}

func NewFintechUseCase(repo domain.FintechRepository) domain.FintechUseCase {
	return &fintechUc{fintechRepo: repo}
}
