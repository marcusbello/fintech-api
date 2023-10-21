package http

import (
	"errors"
	"fintech-api/utils"
)

// PingPong example
type PingPong struct {
	Data string `json:"data"`
} //@name PingPong

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

// Response example
type Response struct {
	Data string `json:"data,omitempty"`
} //@name Response

func (r *Response) Error() string {
	if errors.Is(r, utils.ErrUserNotFound) {
		return utils.ErrUserNotFound.Error()
	}
	return "error:"
}
