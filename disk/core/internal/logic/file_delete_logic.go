package logic

import (
	"context"

	"disk/core/internal/svc"
	"disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteLogic {
	return &FileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDeleteLogic) FileDelete(req *types.FiledeleteRequest) (resp *types.FiledeleteReply, err error) {
	// todo: add your logic here and delete this line

	return
}
