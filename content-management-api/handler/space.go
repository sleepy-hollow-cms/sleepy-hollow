package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"
)

type (
	Space struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
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
	g.PUT("/spaces/:spaceId", r.Update)
	g.GET("/spaces", r.Get)
	g.POST("/spaces", r.Register)
}

func (r *SpaceResource) Get(c echo.Context) error {
	spaces, err := r.SpaceUseCase.Find()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	resultSpaces := make([]Space, len(spaces))
	for i, space := range spaces {
		resultSpaces[i] = Space{
			ID:        space.ID.String(),
			Name:      space.Name.String(),
			CreatedAt: space.CreatedAt.Time().Format(time.RFC3339),
			UpdatedAt: space.UpdatedAt.Time().Format(time.RFC3339),
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
		ID:        space.ID.String(),
		Name:      space.Name.String(),
		CreatedAt: space.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt: space.UpdatedAt.Time().Format(time.RFC3339),
	})
}

func (r *SpaceResource) Update(c echo.Context) error {
	s := Space{}

	if err := c.Bind(&s); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	now := time.Now()
	space, err := r.SpaceUseCase.Register(domain.Space{
		Name:      domain.Name(s.Name),
		UpdatedAt: domain.UpdatedAt(now),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &Space{
		ID:        space.ID.String(),
		Name:      space.Name.String(),
		CreatedAt: space.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt: space.UpdatedAt.Time().Format(time.RFC3339),
	})
}

func (r *SpaceResource) Register(c echo.Context) error {
	s := Space{}

	if err := c.Bind(&s); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	now := time.Now()
	space, err := r.SpaceUseCase.Register(domain.Space{
		Name:      domain.Name(s.Name),
		CreatedAt: domain.CreatedAt(now),
		UpdatedAt: domain.UpdatedAt(now),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &Space{
		ID:        space.ID.String(),
		Name:      space.Name.String(),
		CreatedAt: space.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt: space.UpdatedAt.Time().Format(time.RFC3339),
	})
}
