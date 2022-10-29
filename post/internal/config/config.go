package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf

	App struct {
		MaxCommentCount       int
		AttachmentIncomeRate  float64
		DefaultContextTimeout int
		DefaultPageSize       int
		MaxPageSize           int
	}
}
