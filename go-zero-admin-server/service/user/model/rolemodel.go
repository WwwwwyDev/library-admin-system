package model

import (
	"gorm.io/gorm"
)

type (
	RoleModel interface {
		AddRole(role *Role) (bool,error)
		UpdateRole(role *Role) (bool,error)
		DelRole(role *Role) (bool,error)
		DelRoleUserByUserId(userId uint) (error)
		GetAllRoles()([]Role,error)
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


func (d defaultRoleModel) AddRole(role *Role) (bool, error) {
	err := d.conn.Model(role).Create(role).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultRoleModel) UpdateRole(role *Role) (bool, error) {
	err := d.conn.Model(role).Save(role).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultRoleModel) DelRole(role *Role) (bool, error) {
	err := d.conn.Model(role).Delete(role).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultRoleModel) DelRoleUserByUserId(userId uint) (error){
	err := d.conn.Exec("DELETE FROM user_role WHERE user_id = ?", userId).Error
	if err != nil{
		return err
	}
	return nil
}
func (d defaultRoleModel) GetAllRoles() ([]Role, error) {
	var roles []Role
	err := d.conn.Model(&Role{}).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}


