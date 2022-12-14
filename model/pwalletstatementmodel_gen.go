// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	pWalletStatementFieldNames          = builder.RawFieldNames(&PWalletStatement{})
	pWalletStatementRows                = strings.Join(pWalletStatementFieldNames, ",")
	pWalletStatementRowsExpectAutoSet   = strings.Join(stringx.Remove(pWalletStatementFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")
	pWalletStatementRowsWithPlaceHolder = strings.Join(stringx.Remove(pWalletStatementFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), "=?,") + "=?"

	cachePaopaoPWalletStatementIdPrefix = "cache:paopao:pWalletStatement:id:"
)

type (
	pWalletStatementModel interface {
		Insert(ctx context.Context, data *PWalletStatement) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PWalletStatement, error)
		Update(ctx context.Context, data *PWalletStatement) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPWalletStatementModel struct {
		sqlc.CachedConn
		table string
	}

	PWalletStatement struct {
		Id              int64  `db:"id"`               // 账单ID
		UserId          int64  `db:"user_id"`          // 用户ID
		ChangeAmount    int64  `db:"change_amount"`    // 变动金额
		BalanceSnapshot int64  `db:"balance_snapshot"` // 资金快照
		Reason          string `db:"reason"`           // 变动原因
		PostId          int64  `db:"post_id"`          // 关联动态
		CreatedOn       int64  `db:"created_on"`       // 创建时间
		ModifiedOn      int64  `db:"modified_on"`      // 修改时间
		DeletedOn       int64  `db:"deleted_on"`       // 删除时间
		IsDel           int64  `db:"is_del"`           // 是否删除 0 为未删除、1 为已删除
	}
)

func newPWalletStatementModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultPWalletStatementModel {
	return &defaultPWalletStatementModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`p_wallet_statement`",
	}
}

func (m *defaultPWalletStatementModel) Delete(ctx context.Context, id int64) error {
	paopaoPWalletStatementIdKey := fmt.Sprintf("%s%v", cachePaopaoPWalletStatementIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, paopaoPWalletStatementIdKey)
	return err
}

func (m *defaultPWalletStatementModel) FindOne(ctx context.Context, id int64) (*PWalletStatement, error) {
	paopaoPWalletStatementIdKey := fmt.Sprintf("%s%v", cachePaopaoPWalletStatementIdPrefix, id)
	var resp PWalletStatement
	err := m.QueryRowCtx(ctx, &resp, paopaoPWalletStatementIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pWalletStatementRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPWalletStatementModel) Insert(ctx context.Context, data *PWalletStatement) (sql.Result, error) {
	paopaoPWalletStatementIdKey := fmt.Sprintf("%s%v", cachePaopaoPWalletStatementIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pWalletStatementRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ChangeAmount, data.BalanceSnapshot, data.Reason, data.PostId, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel)
	}, paopaoPWalletStatementIdKey)
	return ret, err
}

func (m *defaultPWalletStatementModel) Update(ctx context.Context, data *PWalletStatement) error {
	paopaoPWalletStatementIdKey := fmt.Sprintf("%s%v", cachePaopaoPWalletStatementIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pWalletStatementRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ChangeAmount, data.BalanceSnapshot, data.Reason, data.PostId, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel, data.Id)
	}, paopaoPWalletStatementIdKey)
	return err
}

func (m *defaultPWalletStatementModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePaopaoPWalletStatementIdPrefix, primary)
}

func (m *defaultPWalletStatementModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pWalletStatementRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPWalletStatementModel) tableName() string {
	return m.table
}
