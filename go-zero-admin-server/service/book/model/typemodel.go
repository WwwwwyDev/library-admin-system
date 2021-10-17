package model

import (
	"gorm.io/gorm"
)

type (
	TypeModel interface {
		GetTypeByNameLike(name string) ([]Type, error)
		GetTypeById(id uint) (*Type, error)
		GetTypes(page int, limit int, typeS *Type) ([]Type, int64, error)
		IsExistTypeById(id uint) (bool, error)
		IsExistTypeByName(name string) (bool, error)
		AddType(_type *Type) (bool, error)
		UpdateType(_type *Type) (bool, error)
		DelType(_type *Type) (bool, error)
	}

	defaultTypeModel struct {
		conn  *gorm.DB
	}

	Type struct {
		gorm.Model
		Name  string `json:"name" gorm:"type:varchar(191);comment:种类名;not null"` //种类名
		Intro string `json:"info"  gorm:"type:varchar(1000);comment:种类介绍"`       //种类介绍
		Book   []Book `json:"books"`
	}
)


func NewTypeModel(conn *gorm.DB) TypeModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Type{})
	return &defaultTypeModel{
		conn:  conn.Debug(),
	}
}

func (d defaultTypeModel) GetTypeByNameLike(name string) ([]Type, error) {
	var _types []Type
	err := d.conn.Model(&Type{}).Where("name like ?", "%"+name+"%").Find(&_types).Error
	if err != nil {
		return nil, err
	}
	return _types, nil
}

func (d defaultTypeModel) GetTypeById(id uint) (*Type, error) {
	var _type *Type
	err := d.conn.Model(&Type{}).Preload("Book").Where("id = ?", id).First(&_type).Error
	if err != nil {
		return nil, err
	}
	return _type, nil
}

func (d defaultTypeModel) IsExistTypeById(id uint) (bool, error){
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

func (d defaultTypeModel) IsExistTypeByName(name string) (bool, error){
	var total int64
	err := d.conn.Model(&Type{}).Where("name = ?", name).Count(&total).Error
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

func (d defaultTypeModel) AddType(_type *Type) (bool, error) {
	err := d.conn.Model(_type).Create(_type).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultTypeModel) UpdateType(_type *Type) (bool, error) {
	err := d.conn.Model(_type).Save(_type).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultTypeModel) DelType(_type *Type) (bool, error) {
	err := d.conn.Model(_type).Delete(_type).Error
	if err != nil {
		return false,err
	}
	return true,nil
}
