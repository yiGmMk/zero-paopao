package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PUserModel = (*customPUserModel)(nil)

type (
	// PUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPUserModel.
	PUserModel interface {
		pUserModel
		Fetch(ctx context.Context, rowBuilder squirrel.SelectBuilder, offset, limit uint64) ([]*PUser, error)
		RowBuilder() squirrel.SelectBuilder
	}

	customPUserModel struct {
		*defaultPUserModel
	}
)

// NewPUserModel returns a model for the database table.
func NewPUserModel(conn sqlx.SqlConn, cache cache.CacheConf) PUserModel {
	return &customPUserModel{
		defaultPUserModel: newPUserModel(conn, cache),
	}
}

func (m *defaultPUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(pUserRows).From(m.table)
}

func (m *defaultPUserModel) Fetch(ctx context.Context, rowBuilder squirrel.SelectBuilder, offset, limit uint64) (
	[]*PUser, error) {
	if offset >= 0 && limit > 0 {
		rowBuilder = rowBuilder.Offset(offset).Limit(limit)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "user sql")
	}

	var resp []*PUser
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, errors.Wrapf(err, "query user:%s", query)
	}
}
