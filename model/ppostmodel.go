package model

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/Masterminds/squirrel"
)

var _ PPostModel = (*customPPostModel)(nil)

type (
	// PPostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPPostModel.
	PPostModel interface {
		pPostModel
		Fetch(ctx context.Context, rowBuilder squirrel.SelectBuilder, offset, limit uint64) ([]*PPost, error)
		RowBuilder() squirrel.SelectBuilder
	}

	customPPostModel struct {
		*defaultPPostModel
	}
)

// NewPPostModel returns a model for the database table.
func NewPPostModel(conn sqlx.SqlConn) PPostModel {
	return &customPPostModel{
		defaultPPostModel: newPPostModel(conn),
	}
}

func (m *defaultPPostModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(pPostRows).From(m.table)
}

func (m *defaultPPostModel) Fetch(ctx context.Context, rowBuilder squirrel.SelectBuilder, offset, limit uint64) ([]*PPost, error) {
	rowBuilder = rowBuilder.OrderBy("is_top DESC, latest_replied_on DESC")

	if offset >= 0 && limit > 0 {
		rowBuilder = rowBuilder.Offset(offset).Limit(limit)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "post sql")
	}

	var resp []*PPost
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, errors.Wrapf(err, "query post:%s", query)
	}
}
