package services

import (
	"errors"

	"github.com/vpiyush/go-kit-sample/pkg/storage/inmem"
)

type PairService interface {
	Insert(key string, value string) (*inmem.Pair, error)
	Get(key string) (*inmem.Pair, error)
}

type Repository interface {
	Insert(key string, value string) (*inmem.Pair, error)
	Get(key string) (*inmem.Pair, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Insert(key, value string) (*inmem.Pair, error) {
	_, err := s.r.Get(key)
	// if there is no error, we can assme key alredy exists
	if err == nil {
		return nil, errors.New("key already exists")
	}
	// extra validation should be added here if anything needed
	// after queriying from database
	return s.r.Insert(key, value)
}

func (s *service) Get(key string) (*inmem.Pair, error) {
	return s.r.Get(key)
}
