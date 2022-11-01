// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	schemaMigrationsFieldNames          = builder.RawFieldNames(&SchemaMigrations{})
	schemaMigrationsRows                = strings.Join(schemaMigrationsFieldNames, ",")
	schemaMigrationsRowsExpectAutoSet   = strings.Join(stringx.Remove(schemaMigrationsFieldNames, "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	schemaMigrationsRowsWithPlaceHolder = strings.Join(stringx.Remove(schemaMigrationsFieldNames, "`version`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
)

type (
	schemaMigrationsModel interface {
		Insert(ctx context.Context, data *SchemaMigrations) (sql.Result, error)
		FindOne(ctx context.Context, version int64) (*SchemaMigrations, error)
		Update(ctx context.Context, data *SchemaMigrations) error
		Delete(ctx context.Context, version int64) error
	}

	defaultSchemaMigrationsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SchemaMigrations struct {
		Version int64 `db:"version"`
		Dirty   int64 `db:"dirty"`
	}
)

func newSchemaMigrationsModel(conn sqlx.SqlConn) *defaultSchemaMigrationsModel {
	return &defaultSchemaMigrationsModel{
		conn:  conn,
		table: "`schema_migrations`",
	}
}

func (m *defaultSchemaMigrationsModel) Delete(ctx context.Context, version int64) error {
	query := fmt.Sprintf("delete from %s where `version` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, version)
	return err
}

func (m *defaultSchemaMigrationsModel) FindOne(ctx context.Context, version int64) (*SchemaMigrations, error) {
	query := fmt.Sprintf("select %s from %s where `version` = ? limit 1", schemaMigrationsRows, m.table)
	var resp SchemaMigrations
	err := m.conn.QueryRowCtx(ctx, &resp, query, version)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSchemaMigrationsModel) Insert(ctx context.Context, data *SchemaMigrations) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, schemaMigrationsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Version, data.Dirty)
	return ret, err
}

func (m *defaultSchemaMigrationsModel) Update(ctx context.Context, data *SchemaMigrations) error {
	query := fmt.Sprintf("update %s set %s where `version` = ?", m.table, schemaMigrationsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Dirty, data.Version)
	return err
}

func (m *defaultSchemaMigrationsModel) tableName() string {
	return m.table
}