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

	apiRoutes := r.Group("/api")
	{
		apiRoutes.POST("/register", handler.RegisterHandler)
		apiRoutes.POST("/signin", handler.LoginHandler)
	}

	userProtectedRoutes := apiRoutes.Group("/:username", AuthorizeJWT())
	{
		userProtectedRoutes.GET("", handler.GetUserHandler)
		userProtectedRoutes.GET("/account", handler.GetAccountHandler)
		userProtectedRoutes.POST("/transfer", handler.TransferMoneyHandler)
	}

	//r.GET("/me", handler.GetUserHandler)
}

func (h FintechHandler) LoginHandler(c *gin.Context) {
	//get details
	var req domain.LoginRequest
	err := c.ShouldBind(&req)
	//clean inputs
	user := strings.ToLower(req.UserName)
	err = h.fintechUc.LoginUc(c, user, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err})
		return
	}
	// generate and add token to header
	jwtToken, err := utils.GenerateToken(user)
	c.Header("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	c.JSON(http.StatusOK, gin.H{"Data": "successfully logged in as"})
}

func (h FintechHandler) RegisterHandler(c *gin.Context) {
	var req domain.RegisterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"Error": err})
		return
	}
	log.Println("Success on binding")
	user := strings.ToLower(req.UserName)
	resp, err := h.fintechUc.RegisterUserUc(c, user, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": fmt.Sprintf("Successfully signed up with %s as your name, visit login page", resp)})
}

func (h FintechHandler) GetUserHandler(c *gin.Context) {
	getUserName := c.Param("username")
	resp, err := h.fintechUc.GetUserUc(c, getUserName)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h FintechHandler) GetAccountHandler(c *gin.Context) {
	getUserName := c.Param("username")
	resp, err := h.fintechUc.GetAccountUc(c, getUserName)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h FintechHandler) TransferMoneyHandler(c *gin.Context) {
	getUserName := c.Param("username")
	var req domain.TransferRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Error binding")
		c.JSON(http.StatusBadRequest, gin.H{"Data": err})
		return
	}
	var resp domain.Account
	if getUserName != "" {
		resp, err = h.fintechUc.TransferMoneyUc(c, req.From, req.To, req.Amount)
		if err != nil {
			log.Printf("Error in usecase")
			c.JSON(http.StatusBadRequest, gin.H{"Data": err})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}
