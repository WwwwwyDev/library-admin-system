package vip

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-admin-server/api/internal/logic/vip"
	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"
)

func AddVipHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddVipReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := vip.NewAddVipLogic(r.Context(), ctx)
		resp, err := l.AddVip(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
