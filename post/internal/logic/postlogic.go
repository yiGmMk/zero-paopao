package logic

import (
	"context"

	"github.com/yiGmMk/zero-paopao/post/internal/svc"
	"github.com/yiGmMk/zero-paopao/post/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostLogic {
	return PostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostLogic) Post(req types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
