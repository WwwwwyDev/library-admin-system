package jwt

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-admin-server/api/internal/logic/jwt"
	"go-zero-admin-server/api/internal/svc"
)

func ChangePasswordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := jwt.NewChangePasswordLogic(r.Context(), ctx)
		resp, err := l.ChangePassword()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
