package service

import (
	"page-pal/model"
	"page-pal/util"
)

func AllUsers() (*[]model.User, error) {
	db := util.GetDb()

  var users []model.User
  res := db.Find(&users)
  
  if res.Error != nil {
    return nil, res.Error
  }

  return &users, nil
}
