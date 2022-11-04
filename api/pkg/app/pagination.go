package app

import (
	"github.com/yiGmMk/zero-paopao/api/internal/config"
)

func GetPageOffset(page, pageSize int, config *config.Config) (offset, limit int) {
	limit = pageSize
	if limit <= 0 {
		limit = config.App.DefaultPageSize
	} else if limit > config.App.MaxPageSize {
		limit = config.App.MaxPageSize
	}
	offset = (page - 1) * limit
	return
}
