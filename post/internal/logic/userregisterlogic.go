package logic

import (
	"context"
	"regexp"
	"unicode/utf8"

	"github.com/pkg/errors"
	"github.com/yiGmMk/zero-paopao/post/internal/svc"
	"github.com/yiGmMk/zero-paopao/post/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ValidUsername 验证用户
func ValidUsername(username string) error {
	// 检测用户是否合规
	if utf8.RuneCountInString(username) < 3 || utf8.RuneCountInString(username) > 12 {
		return errors.Errorf("errcode.UsernameLengthLimit")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
		return errors.Errorf("errcode.UsernameCharLimit")
	}

	// 重复检查
	// user, _ := ds.GetUserByUsername(username)

	// if user.Model != nil && user.ID > 0 {
	// 	return errcode.UsernameHasExisted
	// }

	return nil
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (
	resp *types.UserRegisterRes, err error) {

	resp = &types.UserRegisterRes{}

	return
}
