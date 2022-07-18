package core

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlGorm(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		os.Exit(0)
	} else {
		fmt.Println("mysql " + dsn + " connect successfully")
	}
	return db
}
