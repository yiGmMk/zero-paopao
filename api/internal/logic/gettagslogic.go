package logic

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/yiGmMk/zero-paopao/api/internal/datatypes"
	"github.com/yiGmMk/zero-paopao/api/internal/svc"
	"github.com/yiGmMk/zero-paopao/api/internal/types"
	"github.com/yiGmMk/zero-paopao/api/pkg/app"
	"github.com/yiGmMk/zero-paopao/model"

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
	offset, limit := app.GetPageOffset(1, req.Num, &l.svcCtx.Config)
	b := l.svcCtx.TagModel.RowBuilder()

	// 热门标签
	if req.Type == string(datatypes.TagTypeHot) {
		b.OrderBy("quote_num DESC")
	}
	// 新标签
	if req.Type == string(datatypes.TagTypeNew) {
		b.OrderBy("id DESC")
	}
	tags, err := l.svcCtx.TagModel.Fetch(l.ctx, b, uint64(offset), uint64(limit))
	if err != nil {
		return nil, err
	}

	var (
		userIds []int64
		users   []*model.PUser
		mUsers  map[int64]*model.PUser = make(map[int64]*model.PUser)
	)

	for _, v := range tags {
		userIds = append(userIds, v.UserId)
	}

	if len(userIds) > 0 {
		b := l.svcCtx.UserModel.RowBuilder().Where(squirrel.Eq{"id": userIds})
		users, err = l.svcCtx.UserModel.Fetch(l.ctx, b, 0, 0)
		if err != nil {
			return nil, err
		}
		for _, v := range users {
			mUsers[v.Id] = v
		}
	}

	data := []*types.TagFormated{}
	for _, v := range tags {
		format := types.TagFormated{}
		_ = copier.Copy(&format, v)
		if u, ok := mUsers[format.UserID]; ok {
			uFormat := types.UserFormated{}
			_ = copier.Copy(&uFormat, u)
		}
		data = append(data, &format)
	}
	resp = &types.GetTagResp{
		Data: data,
	}
	return
}
