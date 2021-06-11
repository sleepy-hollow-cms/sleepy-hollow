package port

import "content-management-api/domain"

type Space interface {
	FindByID(id domain.SpaceID) (domain.Space, error)
}
