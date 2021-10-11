package model

import (
	"gorm.io/gorm"
)

type (
	UserModel interface {
		GetUserByUsername(username string) (*User, error)
		GetUserByUsernameLike(username string) ([]User, error)
		GetUserById(id uint) (*User, error)
		GetUsers(page int, limit int, userS *User) ([]User, int64, error)
		IsExistUserById(id uint) (bool, error)
		IsExistUserByUsername(username string) (bool, error)
		AddUser(user *User) (bool,error)
		UpdateUser(user *User) (bool,error)
		DelUser(user *User) (bool,error)
		DelSomeUser(id []uint) (bool,error)
	}

	defaultUserModel struct {
		conn  *gorm.DB
	}

	User struct {
		gorm.Model
		Username       string   `json:"username" gorm:"type:varchar(191);comment:用户名;not null"` //用户名
		Password      string   `json:"password"  gorm:"type:varchar(191);comment:密码md5值;not null"`       //密码md5值
		Salt     string   `json:"salt" gorm:"type:varchar(191);comment:密码盐"`       //密码盐
		Info       string   `json:"info"  gorm:"type:varchar(1000);comment:用户信息"`       //用户信息
		Role   []Role `json:"role" gorm:"many2many:user_role"`
	}
)


func NewUserModel(conn *gorm.DB) UserModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&User{})
	return &defaultUserModel{
		conn:  conn,
	}
}

func (d defaultUserModel) GetUserByUsernameLike(username string) ([]User, error) {
	var users []User
	err := d.conn.Model(&User{}).Where("username like ?", "%"+username+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (d defaultUserModel) GetUserByUsername(username string) (*User, error) {
	var user *User
	err := d.conn.Model(&User{}).Preload("Role").Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d defaultUserModel) GetUserById(id uint) (*User, error) {
	var user *User
	err := d.conn.Model(&User{}).Preload("Role").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d defaultUserModel) IsExistUserById(id uint) (bool, error){
	var total int64
	err := d.conn.Model(&User{}).Where("id = ?", id).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 0{
		return false, nil
	}
	return true,nil
}

func (d defaultUserModel) IsExistUserByUsername(username string) (bool, error){
	var total int64
	err := d.conn.Model(&User{}).Where("username = ?", username).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 0{
		return false, nil
	}
	return true,nil
}

func (d defaultUserModel) GetUsers(page int, limit int, userS *User) ([]User, int64, error) {
	var users []User
	var total int64
	temp := d.conn.Model(&User{}).Order("id ASC")
	if userS.Username != ""{
		temp = temp.Where("username like ?","%"+userS.Username+"%")
	}
	err := temp.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}


func (d defaultUserModel)  AddUser(user *User) (bool,error){
	err := d.conn.Model(user).Create(user).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultUserModel)  UpdateUser(user *User) (bool,error){
	err := d.conn.Model(user).Save(user).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultUserModel) 	DelUser(user *User) (bool,error){
	err := d.conn.Model(user).Delete(user).Error
	if err != nil {
		return false,err
	}
	return true,nil
}


func (d defaultUserModel)  DelSomeUser(id []uint) (bool,error){
	err := d.conn.Model(&User{}).Delete(&User{}, id).Error
	if err != nil {
		return false,err
	}
	return true,nil
}