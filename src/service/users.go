package service

import (
	"page-pal/model"
	"page-pal/dto"
	"page-pal/util"
)

func AllUsers() (*[]dto.UserResponseDto, error) {
	db := util.GetDb()

  var users []dto.UserResponseDto

	db.Model(&model.User{}).
		Select("users.id user_id, users.name as user_name, count(*) as total_books, sum(books.page_count) as total_pages").
		Joins("inner join read_histories on read_histories.user_id = users.id").
		Joins("inner join books on books.id = read_histories.book_id").
    Group("users.id").
    Scan(&users)

	return &users, nil
}
