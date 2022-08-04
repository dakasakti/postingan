package services

import "github.com/dakasakti/postingan/entities"

type PostTypeService interface {
	Register(user_id uint, data entities.PostTypeRequest) error
	GetAll() ([]entities.PostType, error)
	CheckParamId(id string) (uint, error)
	GetById(id uint) (entities.PostType, error)
	UpdateById(id uint, data entities.PostTypeUpdateRequest) error
	DeleteById(id uint) error
}
