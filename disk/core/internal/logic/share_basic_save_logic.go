package logic

import (
	"context"
	"disk/sql"
	"disk/tool"

	"disk/core/internal/svc"
	"disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, get string) (resp *types.ShareBasicSaveReply, err error) {
	// todo: add your logic here and delete this line
	pool := sql.RepositoryPool{}
	pool.Identity = req.RepositoryIdentity
	err = pool.Find()
	if err != nil {
		return nil, err
	}
	basic := sql.UserRepository{
		Identity:           tool.UUID(),
		UserIdentity:       get,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                pool.Ext,
		Name:               pool.Name,
	}
	basic.Insert()
	resp = new(types.ShareBasicSaveReply)
	resp.Identity = basic.Identity
	return
}
