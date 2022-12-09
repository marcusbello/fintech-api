package http

import (
	"fintech-api/docs"
	"fintech-api/pkg/domain"
	"fintech-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"strings"
)

type FintechHandler struct {
	fintechUc domain.FintechUseCase
}

// NewFintechHandler godoc
// @title          Fintech API
// @version        1.0
// @description    This is a fintech webserver.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:3030
// @BasePath /api/v1

func NewFintechHandler(r *gin.Engine, fintechUC domain.FintechUseCase) {
	handler := FintechHandler{fintechUc: fintechUC}

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")

	apiRoutes := v1.Group("/")
	{
		apiRoutes.POST("/register", handler.RegisterHandler)
		apiRoutes.POST("/signin", handler.LoginHandler)
	}

	userProtectedRoutes := apiRoutes.Group("/user/:username", AuthorizeJWT())
	{
		userProtectedRoutes.GET("", handler.GetUserHandler)
		userProtectedRoutes.GET("/account", handler.GetAccountHandler)
		userProtectedRoutes.POST("/transfer", handler.TransferMoneyHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//r.GET("/me", handler.GetUserHandler)
}

// LoginHandler godoc
// @Summary Customer Login
// @Schemes
// @Description Login endpoint
// @Tags        Authentication
// @Accept		json
// @Produce     json
// @Param       request body domain.LoginRequest true "login details"
// @Success     200    {object} utils.Response
// @Failure		401		{object} utils.Response
// @Router      /signin [post]
func (h FintechHandler) LoginHandler(c *gin.Context) {
	//get details
	var req domain.LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	//clean inputs
	user := strings.ToLower(req.UserName)
	if err = h.fintechUc.LoginUc(c, user, req.Password); err != nil {
		c.Error(err)
		log.Println("LoginHandler: ", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, c.Errors)
		return
	}
	// generate and add token to header
	jwtToken, err := utils.GenerateToken(user)
	if err != nil {
		c.Error(err)
		log.Println("GenerateToken: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Errors)
		return
	}
	c.Header("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("successfully logged in as %s", req.UserName)})
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
	resp, err := h.fintechUc.TransferMoneyUc(c, getUserName, req.To, req.Amount)
	if err != nil {
		log.Printf("Error in usecase")
		c.JSON(http.StatusBadRequest, gin.H{"Data": err})
		return
	}
	c.JSON(http.StatusOK, resp)
}
