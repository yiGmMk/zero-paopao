package tag

import (
	"context"

	"github.com/yiGmMk/zero-paopao/post/internal/svc"
	"github.com/yiGmMk/zero-paopao/post/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagsLogic {
	return &GetTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagsLogic) GetTags(req *types.GetTagReq) (resp *types.GetTagResp, err error) {
	// todo: add your logic here and delete this line

	return
}
