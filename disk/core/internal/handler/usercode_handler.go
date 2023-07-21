package handler

import (
	"net/http"

	"disk/core/internal/logic"
	"disk/core/internal/svc"
	"disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func usercodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Data
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUsercodeLogic(r.Context(), svcCtx)
		resp, err := l.Usercode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
