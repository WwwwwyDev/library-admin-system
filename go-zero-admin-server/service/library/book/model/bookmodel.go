package model

import (
	"gorm.io/gorm"
)

type (
	BookModel interface {
		GetBookByName(name string)  ([]Book, error)
		GetBooks(page int, limit int)  ([]Book, int64 ,error)
	}

	defaultBookModel struct {
		conn *gorm.DB
		table string
	}

	Book struct {
		gorm.Model
		Name  string `json:"Name" gorm:"type:varchar(191);comment:图书名;not null"` //图书名
		Image  string `json:"image"  gorm:"type:varchar(191);comment:图书图片"`      //图书图片
		Author string `json:"author" gorm:"type:varchar(191);comment:图书作者"`    //图书作者
		Info   string `json:"info"  gorm:"type:varchar(191);comment:图书信息"`   //图书信息
	}
)


func NewBookModel(conn *gorm.DB) BookModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(Book{})

	return &defaultBookModel{
		conn : conn,
		table:   "books",
	}
}

func (d defaultBookModel) GetBookByName(name string) ([]Book, error) {
	var book []Book
	err := d.conn.Table(d.table).Where("name like ?", "%"+name+"%").Find(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (d defaultBookModel) GetBooks(page int, limit int)  ([]Book, int64, error) {
	var book []Book
	var total int64
	err := d.conn.Table(d.table).Order("id ASC").Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&book).Error
	if err != nil {
		return nil, 0, err
	}
	return book, total, nil
}