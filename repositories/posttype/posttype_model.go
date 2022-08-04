package repositories

import (
	"github.com/dakasakti/postingan/entities"

	"gorm.io/gorm"
)

type postTypeModel struct {
	Db *gorm.DB
}

func NewPostTypeModel(db *gorm.DB) *postTypeModel {
	return &postTypeModel{Db: db}
}

func (tm *postTypeModel) Insert(data entities.PostType) error {
	err := tm.Db.Create(&data)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (tm *postTypeModel) Gets() ([]entities.PostType, error) {
	var PostTypes []entities.PostType
	err := tm.Db.Preload("User").Find(&PostTypes)
	if err.Error != nil {
		return nil, err.Error
	}

	return PostTypes, nil
}

func (tm *postTypeModel) Get(id uint) (entities.PostType, error) {
	var PostType entities.PostType
	err := tm.Db.Preload("User").First(&PostType, id)
	if err.Error != nil {
		return PostType, err.Error
	}

	return PostType, nil
}

func (tm *postTypeModel) Update(id uint, data entities.PostType) error {
	err := tm.Db.Where("id = ?", id).Updates(data)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (tm *postTypeModel) Delete(id uint) error {
	err := tm.Db.Where("id = ?", id).Delete(&entities.PostType{})
	if err.Error != nil {
		return err.Error
	}

	return nil
}
