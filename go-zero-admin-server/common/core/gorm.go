package core

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewPostgresqlGorm(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		os.Exit(0)
	}else{
		fmt.Println("postgresql "+ dsn +" connect successfully")
	}
	return db
}