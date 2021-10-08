package logic

import (
	"context"
	"fmt"
	"go-zero-admin-server/common/code"
	"go-zero-admin-server/common/errorx"

	"go-zero-admin-server/service/library/book/cmd/api/internal/svc"
	"go-zero-admin-server/service/library/book/cmd/api/internal/types"

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

//  get /book?page=x&limit=x
func (l *GetBooksLogic) GetBooks(req types.GetBooksReq) (*types.Reply, error) {
	page := req.Page
	limit := req.Limit
	if page <=0 || limit <=0{
		return nil, errorx.NewCodeError(code.ParameterError,"参数非法")
	}
	books, total, err := l.svcCtx.BookModel.GetBooks(page, limit)
	fmt.Println(books)
	if err != nil {
		return nil,err
	}
	return &types.Reply{Code: code.Success, Data: map[string]interface{}{"books": books, "total": total}, Msg: "ok"}, nil
}
