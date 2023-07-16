package model

import (
	"github.com/jinzhu/gorm"
)

type Books struct {
	gorm.Model
  ISBN        string `json:"isbn" gorm:"unique"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	PageCount   int `json:"page_count"`
}

type ReadHistories struct {
	gorm.Model
	BookID uint  `gorm:"not null"`
	Book   Books `gorm:"foreignKey:BookID"`
	UserID uint  `gorm:"not null"`
	User   User  `gorm:"foreignKey:UserID"`
	IsRead bool
}
