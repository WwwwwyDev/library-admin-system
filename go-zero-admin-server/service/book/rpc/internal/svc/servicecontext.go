package svc

import (
	"fmt"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/book/model"
	"go-zero-admin-server/service/book/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	BookModel model.BookModel
	TypeModel model.TypeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlcfg := c.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlcfg.User,
		mysqlcfg.Password,
		mysqlcfg.Host,
		mysqlcfg.Port,
		mysqlcfg.Name,
	)
	conn := core.NewMysqlGorm(dsn)
	return &ServiceContext{
		Config:    c,
		BookModel: model.NewBookModel(conn),
		TypeModel: model.NewTypeModel(conn),
	}
}
