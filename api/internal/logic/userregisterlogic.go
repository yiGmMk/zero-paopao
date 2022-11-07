package logic

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/yiGmMk/zero-paopao/api/internal/datatypes"
	"github.com/yiGmMk/zero-paopao/api/internal/svc"
	"github.com/yiGmMk/zero-paopao/api/internal/types"
	"github.com/yiGmMk/zero-paopao/api/pkg/errcode"
	"github.com/yiGmMk/zero-paopao/model"

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

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (
	resp *types.UserRegisterRes, err error) {
	// 用户名校验
	err = ValidUsername(req.Username)
	if err != nil {
		return nil, err
	}
	exist, err := l.svcCtx.CheckIfUsernameExist(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errcode.UsernameHasExisted
	}

	err = CheckPassword(req.Password)
	if err != nil {
		return nil, err
	}

	password, salt := EncryptPasswordAndSalt(req.Password)
	user := &model.PUser{
		Nickname:   req.Username,
		Username:   req.Username,
		Password:   password,
		Salt:       salt,
		Status:     datatypes.UserStatusNormal,
		CreatedOn:  time.Now().Unix(),
		ModifiedOn: time.Now().Unix(),
		Avatar:     GetRandomAvatar(),
	}
	addRes, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}
	if num, err := addRes.RowsAffected(); err == nil && num <= 0 {
		return nil, errors.Errorf("注册失败")
	}
	userid, err := addRes.LastInsertId()
	if err != nil {
		return nil, err
	}
	resp = &types.UserRegisterRes{
		Id:       userid,
		Username: req.Username,
	}
	return
}
