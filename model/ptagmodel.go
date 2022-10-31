package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PTagModel = (*customPTagModel)(nil)

type (
	// PTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPTagModel.
	PTagModel interface {
		pTagModel
		Fetch(ctx context.Context, rowBuilder squirrel.SelectBuilder, offset, limit uint64) ([]*PTag, error)
		RowBuilder() squirrel.SelectBuilder
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

func (m *defaultPTagModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(pTagRows).From(m.table)
}

func (m *defaultPTagModel) Fetch(ctx context.Context, rowBuilder squirrel.SelectBuilder, offset, limit uint64) ([]*PTag, error) {
	if offset >= 0 && limit > 0 {
		rowBuilder = rowBuilder.Offset(offset).Limit(limit)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "tag sql")
	}

	var resp []*PTag
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, errors.Wrapf(err, "query tag:%s", query)
	}
}
