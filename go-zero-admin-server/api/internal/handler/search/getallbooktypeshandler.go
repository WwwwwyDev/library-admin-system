package search

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-admin-server/api/internal/logic/search"
	"go-zero-admin-server/api/internal/svc"
)

func GetAllBookTypesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := search.NewGetAllBookTypesLogic(r.Context(), ctx)
		resp, err := l.GetAllBookTypes()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
