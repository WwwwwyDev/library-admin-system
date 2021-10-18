package vip

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-admin-server/api/internal/logic/vip"
	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"
)

func UpdateVipHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateVipReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := vip.NewUpdateVipLogic(r.Context(), ctx)
		resp, err := l.UpdateVip(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
