package entities

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	PostTypeID  uint           `json:"-" gorm:"not null"`
	PostType    PostType       `json:"post_type" gorm:"foreignkey:PostTypeID"`
	UserID      uint           `json:"-"`
	User        User           `json:"author" gorm:"foreignkey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type PostRequest struct {
	Title       string `json:"title" validate:"required,min=5,max=100"`
	Description string `json:"description" validate:"required,min=5,max=255"`
	PostTypeID  uint   `json:"post_type_id" validate:"required,number"`
}

type PostUpdateRequest struct {
	Title       string `json:"title" validate:"omitempty,min=5,max=100"`
	Description string `json:"description" validate:"omitempty,min=5,max=255"`
	PostTypeID  uint   `json:"post_type_id" validate:"omitempty,number"`
}
