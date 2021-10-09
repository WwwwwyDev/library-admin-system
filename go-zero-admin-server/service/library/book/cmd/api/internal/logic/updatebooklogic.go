package logic

import (
	"context"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/library/book/cmd/api/internal/svc"
	"go-zero-admin-server/service/library/book/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateBookLogic {
	return UpdateBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBookLogic) UpdateBook(req types.UpdateBookReq) (*types.Reply, error) {
	isExist, err := l.svcCtx.BookModel.IsExistBookByID(req.ID)
	if err != nil{
		return nil,err
	}
	if !isExist{
		return nil, errorx.NewCodeError(code.NoFoundError,"图书不存在")
	}
	bookOld, err := l.svcCtx.BookModel.GetBookByID(req.ID)
	if bookOld.Name != req.Name{
		isExist, err = l.svcCtx.BookModel.IsExistBookByName(req.Name)
		if err != nil{
			return nil,err
		}
		if isExist{
			return nil, errorx.NewCodeError(code.RepeatError,"图书名重复")
		}
	}
	bookNew := bookOld
	bookNew.Name = req.Name
	bookNew.Info = req.Info
	bookNew.Author = req.Author
	bookNew.Image = req.Image
	err = l.svcCtx.BookModel.UpdateBook(bookNew)
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success, Msg: "ok"}, nil
}
