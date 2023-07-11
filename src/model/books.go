package model

import "github.com/jinzhu/gorm"

type Books struct {
  gorm.Model
  ISBN string 
  Title string
  Author string
  ImageURL string
  Description string
  PageCount int
}
