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

func IsValidUser(name string, pass string) (bool, error) {
	db := util.GetDb()

  var user model.User
	db.Where("name = ?", name).First(&user)

  return user.Pass == util.Hash(pass), nil
}

func InsertUser(user model.User) error {
	db := util.GetDb()

  user.Pass = util.Hash(user.Pass)
	return db.Create(&user).Error
}
