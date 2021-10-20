package systemlog

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-admin-server/api/internal/logic/systemlog"
	"go-zero-admin-server/api/internal/svc"
)

func DeleteAllSystemlogsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := systemlog.NewDeleteAllSystemlogsLogic(r.Context(), ctx)
		resp, err := l.DeleteAllSystemlogs()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
