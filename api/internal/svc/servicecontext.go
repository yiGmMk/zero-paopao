package svc

import (
	"github.com/yiGmMk/zero-paopao/api/internal/config"
	"github.com/yiGmMk/zero-paopao/model"
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
		PostModel:        model.NewPPostModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		PostContentModel: model.NewPPostContentModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		UserModel:        model.NewPUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		TagModel:         model.NewPTagModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
