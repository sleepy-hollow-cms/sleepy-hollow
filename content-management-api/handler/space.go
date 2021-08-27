package handler

import (
	"content-management-api/domain"
	"content-management-api/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Space struct {
		ID   string `json:"id"`
		Name string `json:"name"`
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
	g.POST("/spaces", r.Register)
	g.GET("/spaces", r.Get)
}

func (r *SpaceResource) Get(c echo.Context) error {
	spaces, err := r.SpaceUseCase.Find()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	resultSpaces := make([]Space, len(spaces))
	for i, space := range spaces {
		resultSpaces[i] = Space{
			ID:   space.ID.String(),
			Name: space.Name.String(),
		}
	}

	return c.JSON(http.StatusOK, resultSpaces)
}

func (r *SpaceResource) GetByID(c echo.Context) error {
	spaceId := c.Param("spaceId")

	space, err := r.SpaceUseCase.FindByID(domain.SpaceID(spaceId))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &Space{
		ID:   space.ID.String(),
		Name: space.Name.String(),
	})
}

func (r *SpaceResource) Register(c echo.Context) error {
	s := Space{}

	if err := c.Bind(&s); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	space, err := r.SpaceUseCase.Register(domain.Space{
		Name: domain.Name(s.Name),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &Space{
		ID:   space.ID.String(),
		Name: space.Name.String(),
	})
}
