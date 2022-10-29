package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PCommentContentModel = (*customPCommentContentModel)(nil)

type (
	// PCommentContentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPCommentContentModel.
	PCommentContentModel interface {
		pCommentContentModel
	}

	customPCommentContentModel struct {
		*defaultPCommentContentModel
	}
)

// NewPCommentContentModel returns a model for the database table.
func NewPCommentContentModel(conn sqlx.SqlConn) PCommentContentModel {
	return &customPCommentContentModel{
		defaultPCommentContentModel: newPCommentContentModel(conn),
	}
}
