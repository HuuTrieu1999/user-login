package http

import (
	"log"
	"login/internal/controller/http/request"
	"login/internal/controller/http/response"
	"login/internal/core/common/utils"
	"login/internal/core/entity/error_code"
	"login/internal/core/port/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	invalidRequestResponse = &response.Response{
		ErrorCode:    error_code.InvalidRequest,
		ErrorMessage: error_code.InvalidRequestErrMsg,
		Status:       false,
	}
)

type UserController struct {
	gin         *gin.Engine
	userService service.UserService
}

func NewUserController(
	gin *gin.Engine,
	userService service.UserService,
) UserController {
	return UserController{
		gin:         gin,
		userService: userService,
	}

}

func (u UserController) InitRouter() {
	api := u.gin.Group("/api")
	api.POST("/login", u.login)
	api.POST("/register", u.register)
}

func (u UserController) login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &invalidRequestResponse)
		return
	}
	resp := u.userService.Login(req)
	log.Printf("login, request %s response %s\n", req.Account, utils.GetJsonString(resp))
	c.JSON(http.StatusOK, resp)
}

func (u UserController) register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &invalidRequestResponse)
		return
	}
	resp := u.userService.Register(req)
	log.Printf("register, request %s response %s\n", req.FullName, utils.GetJsonString(resp))
	c.JSON(http.StatusOK, resp)
}
