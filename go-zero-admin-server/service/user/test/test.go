package main

import (
	"fmt"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/user/model"
	"gorm.io/gorm"
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
	role := model.Role{Model:gorm.Model{ID: 1}}
	userModel.AddUser(&model.User{Username: "232423",Password: "2312313",Role: []model.Role{role,role}})

}
