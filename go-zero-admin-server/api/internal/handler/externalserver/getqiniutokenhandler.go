package externalserver

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-admin-server/api/internal/logic/externalserver"
	"go-zero-admin-server/api/internal/svc"
)

func GetQiniuTokenHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := externalserver.NewGetQiniuTokenLogic(r.Context(), ctx)
		resp, err := l.GetQiniuToken()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
