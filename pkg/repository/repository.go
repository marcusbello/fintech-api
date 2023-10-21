package repository

import (
	"context"
	"fintech-api/pkg/domain"
)

type FintechRepository interface {
	LoginRepository(c context.Context, userName, password string) error
	RegisterUserRepository(c context.Context, userName, email, password string) (string, error)
	GetUserRepository(c context.Context, userName string) (domain.User, error)
	GetAccountRepository(c context.Context, userName string) (domain.Account, error)
	AddMoneyRepository(c context.Context, to string, amount int) (domain.Account, error)
	RemoveMoneyRepository(c context.Context, from string, amount int) (domain.Account, error)
	AddToTransaction(c context.Context, from, to string, amount int) error
}
