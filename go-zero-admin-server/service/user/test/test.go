package main

import (
	"fmt"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/admin/user/model"
)

func main() {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"postgres",
		"wwy123456",
		"admin_user",
		"81.70.8.101",
		5432,
	)
	conn := core.NewPostgresqlGorm(dsn)
	userModel:=model.NewUserModel(conn)
	_ = userModel.DelSomeUser([]uint{1, 3, 4})

}
