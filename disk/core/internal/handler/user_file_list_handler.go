package handler

import (
	"net/http"

	"disk/core/internal/logic"
	"disk/core/internal/svc"
	"disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFilerequset
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		Uid := r.Header.Get("userIdentity")
		l := logic.NewUserFileListLogic(r.Context(), svcCtx)
		resp, err := l.UserFileList(&req, Uid)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
