package logic

import (
	"context"
	"disk/sql"
	"disk/tool"
	"errors"

	"disk/core/internal/svc"
	"disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserloginLogic {
	return &UserloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserloginLogic) Userlogin(req *types.UserRequest) (resp *types.Response, err error) {
	data := sql.UserBasic{Name: req.Name, Password: tool.Md5(req.Password)}
	_, error := data.FindUser()

	if error != nil || data.Id == 0 {
		return &types.Response{
			"",
			"出错了",
		}, errors.New("出错了")
	}
	token, _ := tool.GenerateToken(data.Id, data.Identity, data.Name, 3000)
	return &types.Response{Data: token}, nil
}
