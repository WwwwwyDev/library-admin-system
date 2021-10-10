package model

import (
	"gorm.io/gorm"
)

type (
	RoleModel interface {
		
	}

	defaultRoleModel struct {
		conn  *gorm.DB
	}

	Role struct {
		gorm.Model
		Name  string `json:"name" gorm:"Role:varchar(191);comment:角色名;not null"` //角色名
		Info string `json:"info"  gorm:"Role:varchar(1000);comment:角色信息"`       //角色信息
	}
)


func NewRoleModel(conn *gorm.DB) RoleModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Role{})
	return &defaultRoleModel{
		conn:  conn,
	}
}
