package service

import (
	"github.com/jinzhu/gorm"

	"page-pal/model"
	"page-pal/util"
)

func InsertBook(book model.Books) error {
	db := util.GetDb()

	return db.Create(&book).Error
}

func ExistBook(book model.Books) bool {
	db := util.GetDb()

	var ex_book model.Books
	result := db.Where("isbn = ?", book.ISBN).First(&ex_book)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false
		}

		return false
	}
  
	return true
}

func IsValidBook(book model.Books) bool {
	return (book.PageCount != 0 && book.Title != "" && book.Author != "")
}

func InsertReadHistory(history model.ReadHistories) error {
	db := util.GetDb()

	return db.Create(&history).Error
}
