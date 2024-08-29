package controller

import (
	"github.com/labstack/echo/v4"
	"ims-intro/pkg/controller/request"
	"ims-intro/pkg/controller/response"
	"ims-intro/pkg/service"
	"net/http"
	"time"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService}
}

func (controller *UserController) RegisterUserRoutes(e *echo.Echo) {
	e.POST("/login", controller.Login)
	e.POST("/signup", controller.SignUp)
}

func (controller *UserController) Login(c echo.Context) error {
	var loginRequest request.LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: unable to bind the provided data to the user structure"))
	}

	token, err := controller.userService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (controller *UserController) SignUp(c echo.Context) error {
	var signUpRequest request.SignUpRequest
	err := c.Bind(&signUpRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: unable to bind the provided data to the user structure"))
	}

	err = controller.userService.SignUp(signUpRequest.ToDtoModel())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusCreated)
}
