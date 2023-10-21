package domain

import (
	"context"
)

type FintechUseCase interface {
	LoginUc(c context.Context, userName, password string) error
	RegisterUserUc(c context.Context, userName, email, password string) (string, error)
	GetUserUc(c context.Context, userName string) (User, error)
	GetAccountUc(c context.Context, userName string) (Account, error)
	TransferMoneyUc(c context.Context, from, to string, amount int) (Account, error)
	AddMoneyUc(c context.Context, to string, amount int) (Account, error)
	RemoveMoneyUc(c context.Context, from string, amount int) (Account, error)
}
