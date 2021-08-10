package write

import (
	"content-management-api/domain"
)

type ContentModel struct {
	Name      domain.Name
	Fields    domain.Fields
	CreatedAt domain.CreatedAt
	UpdatedAt domain.UpdatedAt
}
