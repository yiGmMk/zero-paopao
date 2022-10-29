package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PMessageModel = (*customPMessageModel)(nil)

type (
	// PMessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPMessageModel.
	PMessageModel interface {
		pMessageModel
	}

	customPMessageModel struct {
		*defaultPMessageModel
	}
)

// NewPMessageModel returns a model for the database table.
func NewPMessageModel(conn sqlx.SqlConn) PMessageModel {
	return &customPMessageModel{
		defaultPMessageModel: newPMessageModel(conn),
	}
}
