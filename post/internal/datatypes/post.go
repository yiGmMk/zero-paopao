package datatypes

import (
	"strings"

	"github.com/yiGmMk/zero-paopao/model"
	"github.com/yiGmMk/zero-paopao/post/internal/types"
)

func PostFormat(p *model.PPost) *types.PostFormated {
	tagsMap := map[string]int8{}
	for _, tag := range strings.Split(p.Tags, ",") {
		tagsMap[tag] = 1
	}
	return &types.PostFormated{
		ID:              p.Id,
		UserID:          p.UserId,
		User:            &types.UserFormated{},
		Contents:        []*types.PostContentFormated{},
		CommentCount:    p.CommentCount,
		CollectionCount: p.CollectionCount,
		UpvoteCount:     p.UpvoteCount,
		Visibility:      uint8(p.Visibility),
		IsTop:           p.IsTop,
		IsEssence:       p.IsEssence,
		IsLock:          p.IsLock,
		LatestRepliedOn: p.LatestRepliedOn,
		CreatedOn:       p.CreatedOn,
		ModifiedOn:      p.ModifiedOn,
		AttachmentPrice: p.AttachmentPrice,
		Tags:            tagsMap,
		IPLoc:           p.IpLoc,
	}
}

// 类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源

type PostContentT int

const (
	CONTENT_TYPE_TITLE PostContentT = iota + 1
	CONTENT_TYPE_TEXT
	CONTENT_TYPE_IMAGE
	CONTENT_TYPE_VIDEO
	CONTENT_TYPE_AUDIO
	CONTENT_TYPE_LINK
	CONTENT_TYPE_ATTACHMENT
	CONTENT_TYPE_CHARGE_ATTACHMENT
)

var (
	mediaContentType = []PostContentT{
		CONTENT_TYPE_IMAGE,
		CONTENT_TYPE_VIDEO,
		CONTENT_TYPE_AUDIO,
		CONTENT_TYPE_ATTACHMENT,
		CONTENT_TYPE_CHARGE_ATTACHMENT,
	}
)

// PostVisibleT 可访问类型，0公开，1私密，2好友
type PostVisibleT uint8

const (
	PostVisitPublic PostVisibleT = iota
	PostVisitPrivate
	PostVisitFriend
	PostVisitInvalid
)
