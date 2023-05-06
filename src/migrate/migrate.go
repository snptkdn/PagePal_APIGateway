package main

import (
	"page-pal/model"
	"page-pal/util"
)

func main() {
  db := util.GetDb()

  db.AutoMigrate(&model.User{})
}
