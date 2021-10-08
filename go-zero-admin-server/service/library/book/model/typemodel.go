package model

import (
	"gorm.io/gorm"
)

type (
	TypeModel interface {
		GetTypeByName(name string) ([]Type, error)
		GetTypeByID(id uint) (*Type, error)
		GetTypes(page int, limit int, typeS *Type) ([]Type, int64, error)
		IsExistType(id uint) (bool, error)
	}

	defaultTypeModel struct {
		conn  *gorm.DB
	}

	Type struct {
		gorm.Model
		Name  string `json:"name" gorm:"type:varchar(191);comment:种类名;not null"` //种类名
		Intro string `json:"info"  gorm:"type:varchar(1000);comment:种类介绍"`       //种类介绍
		Book   []Book `json:"bookType" gorm:"many2many:book_type"`
	}
)


func NewTypeModel(conn *gorm.DB) TypeModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Type{})
	return &defaultTypeModel{
		conn:  conn,
	}
}

func (d defaultTypeModel) GetTypeByName(name string) ([]Type, error) {
	var _types []Type
	err := d.conn.Model(&Type{}).Preload("Book").Where("name like ?", "%"+name+"%").Find(&_types).Error
	if err != nil {
		return nil, err
	}
	return _types, nil
}

func (d defaultTypeModel) GetTypeByID(id uint) (*Type, error) {
	var _type *Type
	err := d.conn.Model(&Type{}).Preload("Book").Where("id = ?", id).First(&_type).Error
	if err != nil {
		return nil, err
	}
	return _type, nil
}

func (d defaultTypeModel) IsExistType(id uint) (bool, error){
	var total int64
	err := d.conn.Model(&Type{}).Where("id = ?", id).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 0{
		return false, nil
	}
	return true,nil
}

func (d defaultTypeModel) GetTypes(page int, limit int, typeS *Type) ([]Type, int64, error) {
	var _types []Type
	var total int64
	temp := d.conn.Model(&Type{}).Order("id ASC")
	if typeS.Name != "" {
		temp.Where("name like ?", "%"+typeS.Name+"%")
	}
	err := temp.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&_types).Error
	if err != nil {
		return nil, 0, err
	}
	return _types, total, nil
}
