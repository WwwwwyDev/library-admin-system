package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-admin-server/service/library/book/cmd/api/internal/logic"
	"go-zero-admin-server/service/library/book/cmd/api/internal/svc"
	"go-zero-admin-server/service/library/book/cmd/api/internal/types"
)

func deleteBookHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteBookReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteBookLogic(r.Context(), ctx)
		resp, err := l.DeleteBook(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}