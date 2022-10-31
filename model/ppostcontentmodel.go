package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PPostContentModel = (*customPPostContentModel)(nil)

type (
	// PPostContentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPPostContentModel.
	PPostContentModel interface {
		pPostContentModel
		Fetch(ctx context.Context, rowBuilder squirrel.SelectBuilder, offset, limit uint64) ([]*PPostContent, error)
		RowBuilder() squirrel.SelectBuilder
	}

	customPPostContentModel struct {
		*defaultPPostContentModel
	}
)

// NewPPostContentModel returns a model for the database table.
func NewPPostContentModel(conn sqlx.SqlConn) PPostContentModel {
	return &customPPostContentModel{
		defaultPPostContentModel: newPPostContentModel(conn),
	}
}

func (m *defaultPPostContentModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(pPostContentRows).From(m.table)
}

func (m *defaultPPostContentModel) Fetch(ctx context.Context, rowBuilder squirrel.SelectBuilder, offset, limit uint64) (
	[]*PPostContent, error) {

	if offset >= 0 && limit > 0 {
		rowBuilder = rowBuilder.Offset(offset).Limit(limit)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "postcontent sql")
	}

	var resp []*PPostContent
	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, errors.Wrapf(err, "query post_content")
	}
}
