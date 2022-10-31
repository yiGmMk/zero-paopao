package svc

import (
	"github.com/yiGmMk/zero-paopao/model"
	"github.com/yiGmMk/zero-paopao/post/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	PostModel        model.PPostModel
	PostContentModel model.PPostContentModel
	UserModel        model.PUserModel
	TagModel         model.PTagModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		PostModel:        model.NewPPostModel(sqlx.NewMysql(c.Mysql.DataSource)),
		PostContentModel: model.NewPPostContentModel(sqlx.NewMysql(c.Mysql.DataSource)),
		UserModel:        model.NewPUserModel(sqlx.NewMysql(c.Mysql.DataSource)),
		TagModel:         model.NewPTagModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
