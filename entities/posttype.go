package entities

import (
	"time"

	"gorm.io/gorm"
)

type PostType struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Jenis     string         `json:"jenis"`
	Type      string         `json:"type"`
	UserID    uint           `json:"-"`
	User      User           `json:"author" gorm:"foreignkey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type PostTypeRequest struct {
	Jenis string `json:"jenis" validate:"required,min=4,max=50"`
	Type  string `json:"type" validate:"required,min=4,max=50"`
}

type PostTypeUpdateRequest struct {
	Jenis string `json:"jenis" validate:"omitempty,min=4,max=50"`
	Type  string `json:"type" validate:"omitempty,min=4,max=50"`
}
