package repositories

import "github.com/dakasakti/postingan/entities"

type PostTypeModel interface {
	Insert(data entities.PostType) error
	Gets() ([]entities.PostType, error)
	Get(id uint) (entities.PostType, error)
	Update(id uint, data entities.PostType) error
	Delete(id uint) error
}
