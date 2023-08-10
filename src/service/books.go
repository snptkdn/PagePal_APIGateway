package service

import (
	"github.com/jinzhu/gorm"

	"page-pal/model"
	"page-pal/util"
)

func InsertBook(book model.Book) error {
	db := util.GetDb()

	return db.Create(&book).Error
}

func FindBook(isbn string) (*model.Book, error) {
	db := util.GetDb()

  var tar model.Book
  res := db.Where("isbn = ?", isbn).First(&tar)

  if res.Error != nil {
    return nil, res.Error
  } else {
    return &tar, nil
  }
}

func ExistBook(book model.Book) bool {
	db := util.GetDb()

	var ex_book model.Book
	result := db.Where("isbn = ?", book.ISBN).First(&ex_book)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false
		}

		return false
	}
  
	return true
}

func IsValidBook(book model.Book) bool {
	return (book.PageCount != 0 && book.Title != "" && book.Author != "")
}

func InsertReadHistory(history model.ReadHistory) error {
	db := util.GetDb()

	return db.Create(&history).Error
}

func FindHistory(user_id string) ([]model.ReadHistory, error) {
	db := util.GetDb()

  var tar []model.ReadHistory
  db.Preload("Book").Preload("User").Where("user_id = ?", user_id).Find(&tar)

  return tar, nil
}
