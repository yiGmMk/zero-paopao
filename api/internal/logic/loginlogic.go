package logic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/yiGmMk/zero-paopao/api/internal/svc"
	"github.com/yiGmMk/zero-paopao/api/internal/types"
	"github.com/yiGmMk/zero-paopao/api/pkg/errcode"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.AuthReq) (resp *types.AuthResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, errcode.UnauthorizedTokenGenerate
		}
		return nil, err
	}
	if user == nil {
		return nil, errcode.UnauthorizedTokenGenerate
	}

	token, err := l.GenerateToken(user.Id, req.Username)
	if err != nil {
		return nil, err
	}

	resp = &types.AuthResp{
		Token: token,
	}
	return
}

func (l *LoginLogic) GenerateToken(useId int64, userName string) (string, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["uid"] = useId
	claims["username"] = userName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
}
