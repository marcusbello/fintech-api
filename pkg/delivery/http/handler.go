package http

import (
	"fintech-api/pkg/domain"
	"fintech-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type FintechHandler struct {
	fintechUc domain.FintechUseCase
}

func NewFintechHandler(r *gin.Engine, fintechUC domain.FintechUseCase) {
	handler := FintechHandler{fintechUc: fintechUC}

	r.POST("/login", handler.LoginHandler)
	r.POST("/register", handler.RegisterHandler)
	r.GET("/me", handler.GetUserHandler)
}

func (h FintechHandler) LoginHandler(c *gin.Context) {
	var req domain.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"Error": err})
		return
	}
	err = h.fintechUc.LoginUc(c, req.UserName, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err})
		return
	}
	user := strings.ToLower(req.UserName)
	jwtToken, err := utils.GenerateToken(user)
	c.Header("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	c.JSON(http.StatusOK, gin.H{"Data": "successfully logged in"})
}

func (h FintechHandler) RegisterHandler(c *gin.Context) {
	var req domain.RegisterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"Error": err})
		return
	}
	log.Println("Success on binding")
	resp, err := h.fintechUc.RegisterUserUc(c, strings.ToLower(req.UserName), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": fmt.Sprintf("Successfully signed up with %s as your name, visit login page", resp)})
}

func (h FintechHandler) GetUserHandler(c *gin.Context) {
	user := c.GetHeader("Authorization")
	token := strings.Split(user, " ")[1]
	getUserName, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": fmt.Sprintf("%s", getUserName)})
}
