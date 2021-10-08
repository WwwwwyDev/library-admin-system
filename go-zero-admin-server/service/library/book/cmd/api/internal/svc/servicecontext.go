package svc

import (
	"fmt"
	"github.com/tal-tech/go-zero/rest"
	"go-zero-admin-server/common"
	"go-zero-admin-server/service/library/book/cmd/api/internal/config"
	"go-zero-admin-server/service/library/book/cmd/api/internal/middleware"
	"go-zero-admin-server/service/library/book/model"
)

type ServiceContext struct {
	Config    config.Config
	Cors      rest.Middleware
	BookModel model.BookModel
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
	conn := common.NewPostgresqlGorm(dsn)
	return &ServiceContext{
		Config: c,
		BookModel: model.NewBookModel(conn),
		Cors: middleware.NewCorsMiddleware().Handle,
	}
}
