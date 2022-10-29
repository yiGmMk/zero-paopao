package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PCommentModel = (*customPCommentModel)(nil)

type (
	// PCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPCommentModel.
	PCommentModel interface {
		pCommentModel
	}

	customPCommentModel struct {
		*defaultPCommentModel
	}
)

// NewPCommentModel returns a model for the database table.
func NewPCommentModel(conn sqlx.SqlConn) PCommentModel {
	return &customPCommentModel{
		defaultPCommentModel: newPCommentModel(conn),
	}
}
