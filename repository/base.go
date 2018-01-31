package repository

import "go_rtb/internal/model"

type Repository interface {
	FindById(id interface{}) (*model.Model, error)
	Create(*model.Model) error
	Update(*model.Model) error
	Delete(id interface{}) error
}
