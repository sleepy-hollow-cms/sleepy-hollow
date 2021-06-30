package handler

import (
	"content-management-api/domain"
	"content-management-api/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Space struct {
		ID string `json:"id"`
	}
)

type SpaceResource struct {
	SpaceUseCase *usecase.Space
}

func NewSpaceResource(useCase *usecase.Space) *SpaceResource {
	return &SpaceResource{
		SpaceUseCase: useCase,
	}
}

func (r *SpaceResource) Routing(e *echo.Echo) {
	g := e.Group("/v1")
	g.GET("/spaces/:spaceId", r.GetByID)
}

func (r *SpaceResource) GetByID(c echo.Context) error {
	spaceId := c.Param("spaceId")

	space, err := r.SpaceUseCase.FindByID(domain.SpaceID(spaceId))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &Space{
		ID: space.ID.String(),
	})
}
