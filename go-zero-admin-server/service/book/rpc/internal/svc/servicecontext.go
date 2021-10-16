package svc

import (
	"fmt"
	"go-zero-admin-server/common/core"
	"go-zero-admin-server/service/book/model"
	"go-zero-admin-server/service/book/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	BookModel model.BookModel
	TypeModel model.TypeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	psqlcfg := c.Postgresql
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		psqlcfg.User,
		psqlcfg.Password,
		psqlcfg.Name,
		psqlcfg.Host,
		psqlcfg.Port,
	)
	conn := core.NewPostgresqlGorm(dsn)
	return &ServiceContext{
		Config: c,
		BookModel: model.NewBookModel(conn),
		TypeModel: model.NewTypeModel(conn),
	}
}
