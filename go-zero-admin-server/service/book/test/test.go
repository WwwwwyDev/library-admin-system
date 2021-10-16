package main

import (
	"fmt"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/book/model"
)

func main() {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"postgres",
		"wwy123456",
		"library_book",
		"81.70.8.101",
		5432,
	)
	conn := core.NewPostgresqlGorm(dsn)
	bookModel:=model.NewBookModel(conn)
	typeModel:=model.NewTypeModel(conn)
	bookModel.AddBook(&model.Book{Name: "dawd222",Author: "www"})
	fmt.Println(bookModel)
	fmt.Println(typeModel)
}
