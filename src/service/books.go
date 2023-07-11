package service

import (
	"page-pal/model"
	"page-pal/util"
)

func InsertBook(book model.Books) error {
	db := util.GetDb()

  return db.Create(&book).Error
}

