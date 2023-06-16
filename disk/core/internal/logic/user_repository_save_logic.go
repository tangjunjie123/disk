package logic

import (
	"context"
	"disk/sql"
	"disk/tool"

	"disk/core/internal/svc"
	"disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, get string) (resp *types.UserRepositorySaveReply, err error) {
	// todo: add your logic here and delete this line
	ur := sql.UserRepository{
		Identity:           tool.UUID(),
		UserIdentity:       get,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	err = ur.Insert()
	return
}
