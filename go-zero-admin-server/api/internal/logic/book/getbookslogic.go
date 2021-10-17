package book

import (
	"context"
	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"
	"go-zero-admin-server/service/book/rpc/bookclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBooksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBooksLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBooksLogic {
	return GetBooksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBooksLogic) GetBooks(req types.GetBooksReq) (*types.Reply, error) {
	page := req.Page
	limit := req.Limit
	if page <=0 || limit <=0{
		return nil, errorx.NewCodeError(code.ParameterError,"参数非法")
	}
	booksResp, err := l.svcCtx.BookRpc.GetBooks(l.ctx, &bookclient.BooksReq{Page: int64(req.Page), Limit: int64(req.Limit), Name: req.Name,Author: req.Author})
	if err != nil {
		return nil, err
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"books": booksResp.BooksInfo, "total": booksResp.Total}, Msg: "查询图书成功"}, nil
}
