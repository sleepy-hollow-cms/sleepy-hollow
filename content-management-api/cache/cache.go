package cache

import (
	"errors"
	"sync"
)

type Key string

const (
	MongoDB Key = Key("mongodb")
)

type Container struct {
	cache sync.Map
}

type Cache interface {
	Store(key Key, value interface{})
	Load(key Key) (interface{}, error)
}

func NewContainer() Cache {
	return &Container{}
}

func (s *Container) Store(key Key, value interface{}) {
	s.cache.Store(key, value)
}

func (s *Container) Load(key Key) (interface{}, error) {
	v, ok := s.cache.Load(key)
	if !ok {
		return nil, errors.New("not found")
	}
	return v, nil
}
