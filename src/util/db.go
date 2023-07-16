package util

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
  
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 外部キー制約を有効にする
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
