package write

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
)

type ContentModel struct {
	Name   domain.Name
	Fields field.Fields
	CreatedAt domain.CreatedAt
}
