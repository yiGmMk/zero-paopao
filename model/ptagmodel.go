package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PTagModel = (*customPTagModel)(nil)

type (
	// PTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPTagModel.
	PTagModel interface {
		pTagModel
	}

	customPTagModel struct {
		*defaultPTagModel
	}
)

// NewPTagModel returns a model for the database table.
func NewPTagModel(conn sqlx.SqlConn) PTagModel {
	return &customPTagModel{
		defaultPTagModel: newPTagModel(conn),
	}
}
