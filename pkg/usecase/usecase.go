package usecase

import (
	"context"
	"fintech-api/pkg/domain"
	"fintech-api/pkg/repository"
	"math/big"
)

type fintechUc struct {
	fintechRepo repository.FintechRepository
}

func (u fintechUc) AddMoneyUc(c context.Context, to string, amount int) (domain.Account, error) {
	return u.fintechRepo.AddMoneyRepository(c, to, amount)
}

func (u fintechUc) RemoveMoneyUc(c context.Context, from string, amount int) (domain.Account, error) {
	return u.fintechRepo.RemoveMoneyRepository(c, from, amount)
}

func (u fintechUc) LoginUc(c context.Context, userName, password string) error {
	return u.fintechRepo.LoginRepository(c, userName, password)
}

func (u fintechUc) RegisterUserUc(c context.Context, userName, email, password string) (string, error) {
	return u.fintechRepo.RegisterUserRepository(c, userName, email, password)
}

func (u fintechUc) GetUserUc(c context.Context, userName string) (domain.User, error) {
	return u.fintechRepo.GetUserRepository(c, userName)
}

func (u fintechUc) GetAccountUc(c context.Context, userName string) (domain.Account, error) {
	return u.fintechRepo.GetAccountRepository(c, userName)
}

func (u fintechUc) TransferMoneyUc(c context.Context, from, to string, amount int) (domain.Account, error) {
	var Balance big.Int
	return domain.Account{
		//UserName: "",
		Balance: Balance,
	}, nil
}

func NewFintechUseCase(repo repository.FintechRepository) domain.FintechUseCase {
	return &fintechUc{fintechRepo: repo}
}
