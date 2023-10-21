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
	"regexp"
	"strings"
)

type FintechHandler struct {
	fintechUc domain.FintechUseCase
}

// NewFintechHandler godoc
// @title          Fintech Bank API
// @version        1.0.0
// @description    Fintech Bank API, a financial management application written in Go!
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
		apiRoutes.POST("/ping", handler.PingHandler)
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

// PingHandler godoc
// @Summary Ping Healthcheck
// @Schemes
// @Description Ping
// @Tags        HealthCheck
// @Produce     json
// @Success     200 {object} domain.PingPong
// @Router      /ping [get]
func (h *FintechHandler) PingHandler(c *gin.Context) {
	resp := &PingPong{Data: "Pong!"}
	c.JSON(http.StatusOK, resp)
}

// LoginHandler godoc
// @Summary Customer Login
// @Schemes
// @Description Login endpoint
// @ID          Authentication
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body     domain.LoginRequest true "login details"
// @Success     200     {object} utils.Response
// @Failure     401     {object} utils.Response
// @Failure     500     {object} utils.Response
// @Router      /signin [post]
func (h *FintechHandler) LoginHandler(c *gin.Context) {
	//get details
	var req LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		log.Println("LoginHandler 1: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Errors)
		return
	}
	//clean inputs
	user := strings.ToLower(req.UserName)
	if err = h.fintechUc.LoginUc(c, user, req.Password); err != nil {
		c.Error(err)
		log.Println("LoginHandler 2: ", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, c.Errors)
		return
	}
	// generate token
	jwtToken, err := utils.GenerateToken(user)
	if err != nil {
		c.Error(err)
		log.Println("GenerateToken: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Errors)
		return
	}
	//c.Header("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("successfully logged in as %s", req.UserName),
		"accessToken": jwtToken,
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
func (h *FintechHandler) RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	// username and password matching regex rule
	usernameRegex := "^[a-zA-Z0-9_]{4,20}$"
	passwordRegex := "^[a-zA-Z0-9!@#%^&*()_+=-]{8,20}$"
	// compile the regular expression
	usernameRx, err := regexp.Compile(usernameRegex)
	if err != nil {
		// the regular expression is invalid
		c.Error(err) //nolint:errcheck
		log.Println("RegisterHandler: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Errors)
		return
	}
	passwordRx, err := regexp.Compile(passwordRegex)
	if err != nil {
		// the password regular expression is invalid
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// catch binding error
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err) //nolint:errcheck
		log.Println("RegisterHandler: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Errors)
		return
	}
	// check if the username is valid
	if !usernameRx.MatchString(req.UserName) {
		// the username is invalid
		c.Error(err) //nolint:errcheck
		log.Println("RegisterHandler: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Errors)
		return
	}
	// check if password is valid
	if !passwordRx.MatchString(req.Password) {
		// the password is invalid
		c.Error(err) //nolint:errcheck
		log.Println("RegisterHandler: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Errors)
		return
	}

	//log.Println("Success on binding")
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
// @Security bearerAuth
// @Summary  Customer Profile
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
func (h *FintechHandler) GetUserHandler(c *gin.Context) {
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
// @Security bearerAuth
// @Summary  Customer Bank Account
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
func (h *FintechHandler) GetAccountHandler(c *gin.Context) {
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
// @Security bearerAuth
// @Summary  Transfer Money
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
func (h *FintechHandler) TransferMoneyHandler(c *gin.Context) {
	getUserName := c.Param("username")
	var req TransferRequest
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
