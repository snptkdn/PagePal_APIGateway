package service

import (
	"page-pal/model"
	"page-pal/util"
)

func ExistUser(name string) (bool, error) {
	db := util.GetDb()

	var count int
	db.Where("name = ?", name).First(&model.User{}).Count(&count)

	return count > 0, nil
}

func InsertUser(user model.User) error {
	db := util.GetDb()

	return db.Create(&user).Error
}
