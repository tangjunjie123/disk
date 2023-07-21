package logic

import (
	"context"
	"disk/sql"
	"disk/tool"

	"disk/core/internal/svc"
	"disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, get string) (resp *types.ShareBasicCreateReply, err error) {

	basic := sql.ShareBasic{
		Identity:               tool.UUID(),
		ExpiredTime:            req.ExpiredTime,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		UserIdentity:           get,
	}
	err = basic.Insert()
	if err != nil {
		return nil, err
	}
	return &types.ShareBasicCreateReply{
		Identity: basic.Identity,
	}, nil
}
