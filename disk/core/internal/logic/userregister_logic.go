package logic

import (
	"context"
	"disk/core/internal/svc"
	"disk/core/internal/types"
	"disk/sql"
	"disk/tool"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserregisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserregisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserregisterLogic {
	return &UserregisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserregisterLogic) Userregister(req *types.UserRegister) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	basic := sql.UserBasic{Name: req.Name, Password: tool.Md5(req.Password), Identity: tool.UUID(), Email: req.Email}
	code := sql.RedSget(req.Email)
	fmt.Println("cd,Code")
	fmt.Println(code, req.Code)
	if code != req.Code {
		return &types.Response{
			Data: "验证码错误",
			Err:  "",
		}, err
	}
	err = basic.CreateUser()
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Data: "注册成功",
		Err:  "",
	}, err
}
