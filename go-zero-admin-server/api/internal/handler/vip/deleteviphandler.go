package vip

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-admin-server/api/internal/logic/vip"
	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"
)

func DeleteVipHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteVipReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := vip.NewDeleteVipLogic(r.Context(), ctx)
		resp, err := l.DeleteVip(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
