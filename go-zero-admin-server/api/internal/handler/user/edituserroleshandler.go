package user

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-admin-server/api/internal/logic/user"
	"go-zero-admin-server/api/internal/svc"
	"go-zero-admin-server/api/internal/types"
)

func EditUserRolesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditUserRolesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewEditUserRolesLogic(r.Context(), ctx)
		resp, err := l.EditUserRoles(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
