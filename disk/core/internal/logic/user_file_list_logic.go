package logic

import (
	"context"
	"disk/sql"
	"github.com/spf13/viper"

	"disk/core/internal/svc"
	"disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFilerequset, uid string) (resp *types.UserFilereply, err error) {
	// todo: add your logic here and delete this line
	reply := &types.UserFilereply{}
	uf := make([]*types.UserFile, 0)
	v := viper.Viper{}
	if req.Size == 0 {
		req.Size = v.GetInt("PageSize")
	}
	if req.Page == 0 {
		req.Page = 1
	}
	pool := []sql.RepositoryPool{}

	sql.Db.Where("identity = (?)", sql.Db.Table("user_repositories").
		Select("Repository_identity").Where("user_Identity = ?", uid)).Find(&pool).Limit(req.Page).Offset(req.Size)
	for i := 0; i < len(pool); i++ {
		uf = append(uf, &types.UserFile{
			Id:                 int64(pool[i].Id),
			Identity:           pool[i].Identity,
			RepositoryIdentity: uid,
			Name:               pool[i].Name,
			Ext:                pool[i].Ext,
			Path:               pool[i].Path,
			Size:               pool[i].Size,
		})
	}
	reply.List = uf
	reply.Count = int64(len(pool))
	return reply, nil
}
