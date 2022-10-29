package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PCommentReplyModel = (*customPCommentReplyModel)(nil)

type (
	// PCommentReplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPCommentReplyModel.
	PCommentReplyModel interface {
		pCommentReplyModel
	}

	customPCommentReplyModel struct {
		*defaultPCommentReplyModel
	}
)

// NewPCommentReplyModel returns a model for the database table.
func NewPCommentReplyModel(conn sqlx.SqlConn) PCommentReplyModel {
	return &customPCommentReplyModel{
		defaultPCommentReplyModel: newPCommentReplyModel(conn),
	}
}
