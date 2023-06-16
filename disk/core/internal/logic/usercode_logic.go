package logic

import (
	"context"
	"disk/core/internal/svc"
	"disk/core/internal/types"
	"disk/sql"
	"disk/tool"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsercodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsercodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsercodeLogic {
	return &UsercodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsercodeLogic) Usercode(req *types.Data) (resp *types.Data, err error) {
	basic := sql.UserBasic{Email: req.Data}
	findemail, err := basic.Findemail()
	if !findemail {
		return &types.Data{
			Data: "邮箱已存在",
		}, err
	}
	r := tool.RandInt()
	code := strconv.FormatUint(uint64(r), 10)
	fmt.Println(code)
	tool.Mail(code, req.Data)
	sql.RedSadd(req.Data, code, 60)
	return
}
