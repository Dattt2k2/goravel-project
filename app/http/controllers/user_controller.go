package controllers

import (
	"goravel/app/http/requests"
	"goravel/app/models"
	"goravel/app/services/interfaces"

	"github.com/goravel/framework/contracts/http"
)

type UserController struct {
	userService interfaces.UserServiceInterface
}

func NewUserController(userService interfaces.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Index(ctx http.Context) http.Response {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"data": users,
	})
}

func (c *UserController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")
	user, err := c.userService.GetUser(id)
	if err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"error": "User not found",
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"data": user,
	})
}

func (c *UserController) Store(ctx http.Context) http.Response {
	var userRequest requests.UserRequest
	if err := ctx.Request().Bind(&userRequest); err != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"error": err.Error(),
		})
	}

	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	createdUser, err := c.userService.CreateUser(user)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusCreated, http.Json{
		"data": createdUser,
	})
}
