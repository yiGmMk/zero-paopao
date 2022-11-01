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
	pCaptchaFieldNames          = builder.RawFieldNames(&PCaptcha{})
	pCaptchaRows                = strings.Join(pCaptchaFieldNames, ",")
	pCaptchaRowsExpectAutoSet   = strings.Join(stringx.Remove(pCaptchaFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), ",")
	pCaptchaRowsWithPlaceHolder = strings.Join(stringx.Remove(pCaptchaFieldNames, "`id`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`"), "=?,") + "=?"
)

type (
	pCaptchaModel interface {
		Insert(ctx context.Context, data *PCaptcha) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PCaptcha, error)
		Update(ctx context.Context, data *PCaptcha) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPCaptchaModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PCaptcha struct {
		Id         int64  `db:"id"`          // 验证码ID
		Phone      string `db:"phone"`       // 手机号
		Captcha    string `db:"captcha"`     // 验证码
		UseTimes   int64  `db:"use_times"`   // 使用次数
		ExpiredOn  int64  `db:"expired_on"`  // 过期时间
		CreatedOn  int64  `db:"created_on"`  // 创建时间
		ModifiedOn int64  `db:"modified_on"` // 修改时间
		DeletedOn  int64  `db:"deleted_on"`  // 删除时间
		IsDel      int64  `db:"is_del"`      // 是否删除 0 为未删除、1 为已删除
	}
)

func newPCaptchaModel(conn sqlx.SqlConn) *defaultPCaptchaModel {
	return &defaultPCaptchaModel{
		conn:  conn,
		table: "`p_captcha`",
	}
}

func (m *defaultPCaptchaModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPCaptchaModel) FindOne(ctx context.Context, id int64) (*PCaptcha, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pCaptchaRows, m.table)
	var resp PCaptcha
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPCaptchaModel) Insert(ctx context.Context, data *PCaptcha) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, pCaptchaRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Phone, data.Captcha, data.UseTimes, data.ExpiredOn, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel)
	return ret, err
}

func (m *defaultPCaptchaModel) Update(ctx context.Context, data *PCaptcha) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pCaptchaRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Phone, data.Captcha, data.UseTimes, data.ExpiredOn, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel, data.Id)
	return err
}

func (m *defaultPCaptchaModel) tableName() string {
	return m.table
}