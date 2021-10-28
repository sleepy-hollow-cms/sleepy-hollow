package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"
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
	g.DELETE("/user/:userId", u.Delete)
	g.PUT("/user/:userId", u.Update)
}

func (u *UserResource) Register(c echo.Context) error {
	input := UserCreateRequest{}
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

func (u *UserResource) Update(c echo.Context) error {
	userId := c.Param("userId")

	input := UserUpdateRequest{}
	if err := c.Bind(&input); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	updatedUser, err := u.User.Update(
		domain.User{
			Id:   domain.UserId(userId),
			Name: domain.UserName(input.Name),
		},
	)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, User{
		Id:   updatedUser.Id.String(),
		Name: updatedUser.Name.String(),
	})
}

func (u *UserResource) Delete(c echo.Context) error {
	userId := c.Param("userId")

	err := u.User.DeleteById(domain.UserId(userId))

	if err != nil {
		switch err.(type) {
		case usecase.UserNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusNoContent)
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}
