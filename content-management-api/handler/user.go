package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"
	"net/http"
)

type UserResource struct {
	User *usecase.User
}

func NewUserResource(user *usecase.User) *UserResource {
	return &UserResource{
		User: user,
	}
}

func (u *UserResource) Routing(e *echo.Echo) {
	g := e.Group("/v1")
	g.POST("/user", u.Register)
}

func (u *UserResource) Register(c echo.Context) error {
	input := UserRequest{}
	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	user := domain.User{
		Name: domain.UserName(input.Name),
	}

	registeredUser, err := u.User.Register(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &User{
		Id:   registeredUser.Id.String(),
		Name: registeredUser.Name.String(),
	})
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserRequest struct {
	Name string `json:"name"`
}
