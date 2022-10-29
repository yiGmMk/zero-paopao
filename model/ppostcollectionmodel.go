package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PPostCollectionModel = (*customPPostCollectionModel)(nil)

type (
	// PPostCollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPPostCollectionModel.
	PPostCollectionModel interface {
		pPostCollectionModel
	}

	customPPostCollectionModel struct {
		*defaultPPostCollectionModel
	}
)

// NewPPostCollectionModel returns a model for the database table.
func NewPPostCollectionModel(conn sqlx.SqlConn) PPostCollectionModel {
	return &customPPostCollectionModel{
		defaultPPostCollectionModel: newPPostCollectionModel(conn),
	}
}
