package migrate

import (
	"page-pal/model"
	"page-pal/util"
)

func Migrate() {
  db := util.GetDb()

  db.AutoMigrate(&model.User{})
  db.AutoMigrate(&model.Books{})
  db.AutoMigrate(&model.ReadHistories{})
}