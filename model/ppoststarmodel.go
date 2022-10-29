package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PPostStarModel = (*customPPostStarModel)(nil)

type (
	// PPostStarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPPostStarModel.
	PPostStarModel interface {
		pPostStarModel
	}

	customPPostStarModel struct {
		*defaultPPostStarModel
	}
)

// NewPPostStarModel returns a model for the database table.
func NewPPostStarModel(conn sqlx.SqlConn) PPostStarModel {
	return &customPPostStarModel{
		defaultPPostStarModel: newPPostStarModel(conn),
	}
}
