package model

import (
	"gorm.io/gorm"
)

type (
	BookModel interface {
		GetBookByName(name string) ([]Book, error)
		GetBookByID(id uint) (*Book, error)
		GetBooks(page int, limit int, bookS *Book) ([]Book, int64, error)
		IsExistBook(id uint) (bool, error)
		AddBook(book *Book) error
		UpdateBook(book *Book) error
		DelBook(book *Book) error
	}

	defaultBookModel struct {
		conn  *gorm.DB
	}

	Book struct {
		gorm.Model
		Name       string   `json:"name" gorm:"type:varchar(191);comment:图书名;not null"` //图书名
		Image      string   `json:"image"  gorm:"type:varchar(191);comment:图书图片"`       //图书图片
		Author     string   `json:"author" gorm:"type:varchar(191);comment:图书作者"`       //图书作者
		Info       string   `json:"info"  gorm:"type:varchar(1000);comment:图书信息"`       //图书信息
		Type   []Type `json:"bookType" gorm:"many2many:book_type"`
	}
)

func NewBookModel(conn *gorm.DB) BookModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Book{})
	return &defaultBookModel{
		conn:  conn.Debug(),
	}
}

func (d defaultBookModel) GetBookByName(name string) ([]Book, error) {
	var books []Book
	err := d.conn.Model(&Book{}).Preload("Type").Where("name like ?", "%"+name+"%").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (d defaultBookModel) GetBookByID(id uint) (*Book, error) {
	var book *Book
	err := d.conn.Model(&Book{}).Preload("Type").Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (d defaultBookModel) IsExistBook(id uint) (bool, error){
	var total int64
	err := d.conn.Model(&Book{}).Where("id = ?", id).Count(&total).Error
	if err != nil {
		return false, err
	}
	if total == 0{
		return false, nil
	}
	return true,nil
}

func (d defaultBookModel) GetBooks(page int, limit int, bookS *Book) ([]Book, int64, error) {
	var books []Book
	var total int64
	temp := d.conn.Model(&Book{}).Order("id ASC")
	if bookS.Name != ""{
		temp = temp.Where("name like ?","%"+bookS.Name+"%")
	}
	if bookS.Author != ""{
		temp = temp.Where("author like ?", "%"+bookS.Author+"%")
	}
	err := temp.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	return books, total, nil
}


func (d defaultBookModel)  AddBook(book *Book) error{
	err := d.conn.Model(book).Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (d defaultBookModel)  UpdateBook(book *Book) error{
	err := d.conn.Model(book).Save(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (d defaultBookModel) 	DelBook(book *Book) error{
	err := d.conn.Model(book).Delete(book).Error
	if err != nil {
		return err
	}
	return nil
}