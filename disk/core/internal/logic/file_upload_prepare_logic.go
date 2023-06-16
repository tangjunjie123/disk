package logic

import (
	"context"
	"disk/sql"
	"disk/tool"
	"fmt"

	"disk/core/internal/svc"
	"disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareReply, err error) {
	// todo: add your logic here and delete this line
	pool := sql.RepositoryPool{Hash: req.Md5}
	resp = new(types.FileUploadPrepareReply)
	err = pool.FindH()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if pool.Identity != "" {
		resp.Identity = pool.Identity
	} else {
		key, uploadId, err := tool.InitPart(req.Ext)
		if err != nil {
			return nil, err
		}
		resp.Key = key
		resp.UploadId = uploadId
	}
	return
}
