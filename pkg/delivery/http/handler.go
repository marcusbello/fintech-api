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

// @securityDefinitions.apikey bearerAuth
// @in                         header
// @name                       Authorization
// @description                Description for what is this security definition being used

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
// @ID Authentication
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body     domain.LoginRequest true "login details"
// @Success     200     {object} utils.Response
// @Failure     401     {object} utils.Response
// @Failure     500     {object} utils.Response
// @Router      /signin [post]
func (h FintechHandler) LoginHandler(c *gin.Context) {
	//get details
	var req domain.LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("LoginHandler 1: ", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, c.Errors)
		return
	}
	//clean inputs
	user := strings.ToLower(req.UserName)
	if err = h.fintechUc.LoginUc(c, user, req.Password); err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("LoginHandler 2: ", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, c.Errors)
		return
	}
	// generate and add token to header
	jwtToken, err := utils.GenerateToken(user)
	if err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("GenerateToken: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Errors)
		return
	}
	//c.Header("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("successfully logged in as %s", req.UserName),
		"token": jwtToken,
	})
}

// RegisterHandler godoc
// @Summary Customer Register
// @Schemes
// @Description Register endpoint
// @Tags        Sign-Up
// @Accept      json
// @Produce     json
// @Param       request body     domain.RegisterRequest true "login details"
// @Success     200     {object} utils.Response
// @Failure     401     {object} utils.Response
// @Failure     403     {object} utils.Response
// @Failure     500     {object} utils.Response
// @Router      /register [post]
func (h FintechHandler) RegisterHandler(c *gin.Context) {
	var req domain.RegisterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("RegisterHandler: ", err)
		c.AbortWithStatusJSON(http.StatusForbidden, c.Errors)
		return
	}
	log.Println("Success on binding")
	user := strings.ToLower(req.UserName)
	resp, err := h.fintechUc.RegisterUserUc(c, user, req.Email, req.Password)
	if err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("RegisterHandler: ", err)
		c.AbortWithStatusJSON(http.StatusForbidden, c.Errors)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Successfully signed up with %s as your name, visit login page", resp)})
}

// GetUserHandler godoc
// @Security    bearerAuth
// @Summary Customer Profile
// @Schemes
// @Description Customer Profile endpoint
// @Tags        Profile
// @Accept      json
// @Produce     json
// @Param       user_name path     string true "username"
// @Success     200       {object} domain.UserType
// @Failure     401       {object} utils.Response
// @Failure     403       {object} utils.Response
// @Failure     500       {object} utils.Response
// @Router      /user/{user_name} [get]
func (h FintechHandler) GetUserHandler(c *gin.Context) {
	getUserName := c.Param("username")
	resp, err := h.fintechUc.GetUserUc(c, getUserName)
	if err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("GetUserHandler: ", err)
		c.AbortWithStatusJSON(http.StatusForbidden, c.Errors)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAccountHandler godoc
// @Security    bearerAuth
// @Summary Customer Bank Account
// @Schemes
// @Description Customer Account endpoint
// @Tags        Banking
// @Accept      json
// @Produce     json
// @Param       user_name path     string true "username"
// @Success     200       {object} domain.Account
// @Failure     401       {object} utils.Response
// @Failure     403       {object} utils.Response
// @Failure     500       {object} utils.Response
// @Router      /user/{user_name}/account [get]
func (h FintechHandler) GetAccountHandler(c *gin.Context) {
	getUserName := c.Param("username")
	resp, err := h.fintechUc.GetAccountUc(c, getUserName)
	if err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("GetAccountHandler: ", err)
		c.AbortWithStatusJSON(http.StatusForbidden, c.Errors)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// TransferMoneyHandler godoc
// @Security    bearerAuth
// @Summary Transfer Money
// @Schemes
// @Description Customer Transfer endpoint
// @Tags        Banking
// @Accept      json
// @Produce     json
// @Param       user_name path     string                 true "username"
// @Param       request   body     domain.TransferRequest true "login details"
// @Success     200       {object} domain.Account
// @Failure     401       {object} utils.Response
// @Failure     403       {object} utils.Response
// @Failure     500       {object} utils.Response
// @Router      /user/{user_name}/transfer [post]
func (h FintechHandler) TransferMoneyHandler(c *gin.Context) {
	getUserName := c.Param("username")
	var req domain.TransferRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("TransferHandler: ", err)
		c.AbortWithStatusJSON(http.StatusForbidden, c.Errors)
		return
	}
	resp, err := h.fintechUc.TransferMoneyUc(c, getUserName, req.To, req.Amount)
	if err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("TransferMoneyHandler: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Errors)
		return
	}
	c.JSON(http.StatusOK, resp)
}
