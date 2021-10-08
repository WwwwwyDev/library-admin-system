package logic

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/library/book/model"
	"gorm.io/gorm"

	"go-zero-admin-server/service/library/book/cmd/api/internal/svc"
	"go-zero-admin-server/service/library/book/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteBookLogic {
	return DeleteBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBookLogic) DeleteBook(req types.DeleteBookReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.BookModel.IsExistBookByID(req.ID)
	if err != nil{
		return nil,err
	}
	if !isExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"图书不存在")
	}
	err = l.svcCtx.BookModel.DelBook(&model.Book{Model: gorm.Model{ID: req.ID}})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success, Msg: "ok"}, nil
}
