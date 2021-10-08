package main

import (
	"fmt"
	"go-zero-admin-server/common"
	"go-zero-admin-server/service/library/book/model"
)

func main(){

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"postgres",
		"wwy123456",
		"library_book",
		"81.70.8.101",
		5432,
	)
	conn := common.NewPostgresqlGorm(dsn)
	bookModel := model.NewBookModel(conn)
	//typeModel := model.NewTypeModel(conn)

	err := bookModel.AddBook(&model.Book{Author: "779"})
	if err != nil {
		fmt.Println(err)
	}
	//name1, _ := typeModel.GetTypeByID(1)
	//fmt.Println(name1)
}
