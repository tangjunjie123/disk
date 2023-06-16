package handler

import (
	"crypto/md5"
	"disk/sql"
	"fmt"
	"net/http"
	"path"

	"disk/core/internal/logic"
	"disk/core/internal/svc"
	"disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		file, header, err2 := r.FormFile("file")
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		b := make([]byte, header.Size)
		file.Read(b)
		s := fmt.Sprintf("%x", md5.Sum(b))
		pool := sql.RepositoryPool{}
		tx := sql.Db.Where("hash = ?", s).Find(&pool)
		if tx.Error != nil {
			return
		}
		if pool.Hash != "" {
			//httpx.OkJson(w, &pool)
			httpx.OkJsonCtx(r.Context(), w, nil)
			return
		}
		//txoss := tool.FileuploadTxoss(r)
		req.Name = header.Filename
		req.Ext = path.Ext(header.Filename)
		req.Size = header.Size
		req.Hash = s
		req.Path = "txoss"
		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
