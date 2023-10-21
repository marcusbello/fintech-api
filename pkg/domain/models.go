package domain

import (
	"github.com/google/uuid"
	"math/big"
	"time"
)

type User struct {
	UserID    uuid.UUID
	UserName  string
	Email     string
	Password  string
	AddedAt   time.Time
	UpdatedAt time.Time
}

type Account struct {
	AccountID uuid.UUID
	Balance   big.Int
	AddedAt   time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	TransferID uuid.UUID
	From       uuid.UUID
	To         uuid.UUID
	Amount     big.Int
	AddedAt    time.Time
	UpdatedAt  time.Time
}
