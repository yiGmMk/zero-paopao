package logic

import (
	"context"

	"github.com/Masterminds/squirrel"

	"github.com/jinzhu/copier"
	"github.com/yiGmMk/zero-paopao/api/internal/svc"
	"github.com/yiGmMk/zero-paopao/api/internal/types"
	"github.com/yiGmMk/zero-paopao/model"

	"github.com/yiGmMk/zero-paopao/api/internal/datatypes"
	"github.com/yiGmMk/zero-paopao/api/pkg/app"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// simpleCacheIndexGetPosts simpleCacheIndex 专属获取广场推文列表函数
func (l *GetPostsLogic) IndexPosts(req *types.GetPostsReq) (resp *types.GetPostsResp, err error) {
	b := l.svcCtx.PostModel.RowBuilder().
		Where("visibility=?", datatypes.PostVisitPublic)

	offset, limit := app.GetPageOffset(req.Page, req.PageSize, &l.svcCtx.Config)

	data, err := l.svcCtx.PostModel.Fetch(l.ctx, b, uint64(offset), uint64(limit))
	if err != nil {
		return
	}

	formatData, err := l.MergePosts(data)
	if err != nil {
		return
	}

	resp = &types.GetPostsResp{}
	resp.List = formatData
	resp.Pager.Page = req.Page
	resp.Pager.PageSize = req.PageSize
	resp.Pager.TotalRows = len(resp.List)
	return
}

func (l *GetPostsLogic) getPostContentsByIDs(ids []int64) ([]*model.PPostContent, error) {
	b := l.svcCtx.PostContentModel.RowBuilder().
		Where(squirrel.Eq{"post_id": ids}).OrderBy("sort ASC")
	data, err := l.svcCtx.PostContentModel.Fetch(l.ctx, b, 0, 0)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (l *GetPostsLogic) getUsersByIDs(ids []int64) ([]*model.PUser, error) {
	b := l.svcCtx.UserModel.RowBuilder().
		Where(squirrel.Eq{"id": ids})

	data, err := l.svcCtx.UserModel.Fetch(l.ctx, b, 0, 0)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 动态数据组装
func (l *GetPostsLogic) MergePosts(posts []*model.PPost) ([]*types.PostFormated, error) {
	postIds := make([]int64, 0, len(posts))
	userIds := make([]int64, 0, len(posts))
	for _, post := range posts {
		postIds = append(postIds, post.Id)
		userIds = append(userIds, post.UserId)
	}

	postContents, err := l.getPostContentsByIDs(postIds)
	if err != nil {
		return nil, err
	}

	users, err := l.getUsersByIDs(userIds)
	if err != nil {
		return nil, err
	}

	userMap := make(map[int64]*types.UserFormated, len(users))
	for _, user := range users {
		var format types.UserFormated
		_ = copier.Copy(&format, user)
		userMap[user.Id] = &format
	}

	contentMap := make(map[int64][]*types.PostContentFormated, len(postContents))
	for _, content := range postContents {
		var format types.PostContentFormated
		_ = copier.Copy(&format, content)
		contentMap[content.PostId] = append(contentMap[content.PostId], &format)
	}

	// 数据整合
	postsFormated := make([]*types.PostFormated, 0, len(posts))
	for _, post := range posts {
		postFormated := datatypes.PostFormat(post)
		postFormated.User = userMap[post.UserId]
		postFormated.Contents = contentMap[post.Id]
		postsFormated = append(postsFormated, postFormated)
	}
	return postsFormated, nil
}

func (l *GetPostsLogic) GetPosts(req *types.GetPostsReq) (resp *types.GetPostsResp, err error) {
	l.Logger.Infof("%+v", *req)
	//if req.Query == "" && req.Type == "search" {
	resp, err = l.IndexPosts(req)
	if err != nil {
		l.Logger.WithContext(l.ctx).Errorf("get post:%+v", err)
	}
	return
	// } else {
	// 	user, ok := l.ctx.Value("USER").(*model.PUser)
	// 	if !ok {
	// 		err = errors.New("user info  is nil")
	// 		return
	// 	}
	// 	l.Logger.Infof("%+v", user)
	// }
	//return
}
