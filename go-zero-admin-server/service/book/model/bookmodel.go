package model

import (
	"gorm.io/gorm"
)

type (
	BookModel interface {
		GetBookByNameLike(name string) ([]Book, error)
		GetBookByID(id uint) (*Book, error)
		GetBooks(page int, limit int, bookS *Book) ([]Book, int64, error)
		IsExistBookById(id uint) (bool, error)
		IsExistBookByName(name string) (bool, error)
		AddBook(book *Book) (bool, error)
		UpdateBook(book *Book)(bool, error)
		DelBook(book *Book) (bool, error)
		DelSomeBook(ids []uint) (bool, error)
	}

	defaultBookModel struct {
		conn *gorm.DB
	}

	Book struct {
		gorm.Model
		Name   string `json:"name" gorm:"type:varchar(191);comment:图书名;not null"` //图书名
		Image  string `json:"image"  gorm:"type:varchar(191);comment:图书图片"`       //图书图片
		Author string `json:"author" gorm:"type:varchar(191);comment:图书作者"`       //图书作者
		Info   string `json:"info"  gorm:"type:varchar(1000);comment:图书信息"`       //图书信息
		Type   *Type   `json:"type"`
		TypeID uint
	}
)

func NewBookModel(conn *gorm.DB) BookModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Book{})
	return &defaultBookModel{
		conn: conn.Debug(),
	}
}

func (d defaultBookModel) GetBookByNameLike(name string) ([]Book, error) {
	var book []Book
	err := d.conn.Model(&Book{}).Preload("Type").Where("name like ?", "%"+name+"%").Find(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (d defaultBookModel) GetBookByID(id uint) (*Book, error) {
	var book *Book
	err := d.conn.Model(&Book{}).Preload("Type").Where("id = ?", id).Find(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (d defaultBookModel) IsExistBookById(id uint) (bool, error) {
	var total int64
	err := d.conn.Model(&Book{}).Where("id = ?", id).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 0 {
		return false, nil
	}
	return true, nil
}

func (d defaultBookModel) IsExistBookByName(name string) (bool, error) {
	var total int64
	err := d.conn.Model(&Book{}).Where("name = ?", name).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 0 {
		return false, nil
	}
	return true, nil
}

func (d defaultBookModel) GetBooks(page int, limit int, bookS *Book) ([]Book, int64, error) {
	var books []Book
	var total int64
	temp := d.conn.Model(&Book{}).Preload("Type").Order("id ASC")
	if bookS.Name != "" {
		temp = temp.Where("name like ?", "%"+bookS.Name+"%")
	}
	if bookS.Author != "" {
		temp = temp.Where("author like ?", "%"+bookS.Author+"%")
	}
	err := temp.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	return books, total, nil
}

func (d defaultBookModel) AddBook(book *Book) (bool, error) {
	err := d.conn.Model(book).Create(book).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultBookModel) UpdateBook(book *Book) (bool, error) {
	err := d.conn.Model(book).Save(book).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultBookModel) DelBook(book *Book) (bool, error) {
	err := d.conn.Model(book).Delete(book).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

func (d defaultBookModel) DelSomeBook(ids []uint) (bool, error) {
	err := d.conn.Model(&Book{}).Delete(&Book{}, ids).Error
	if err != nil {
		return false, err
	}
	return true, nil
}