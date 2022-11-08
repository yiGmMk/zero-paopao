package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PAttachmentModel = (*customPAttachmentModel)(nil)

type (
	// PAttachmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPAttachmentModel.
	PAttachmentModel interface {
		pAttachmentModel
	}

	customPAttachmentModel struct {
		*defaultPAttachmentModel
	}
)

// NewPAttachmentModel returns a model for the database table.
func NewPAttachmentModel(conn sqlx.SqlConn, cache cache.CacheConf) PAttachmentModel {
	return &customPAttachmentModel{
		defaultPAttachmentModel: newPAttachmentModel(conn, cache),
	}
}
