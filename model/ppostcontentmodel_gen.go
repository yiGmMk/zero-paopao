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
	pPostContentFieldNames          = builder.RawFieldNames(&PPostContent{})
	pPostContentRows                = strings.Join(pPostContentFieldNames, ",")
	pPostContentRowsExpectAutoSet   = strings.Join(stringx.Remove(pPostContentFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	pPostContentRowsWithPlaceHolder = strings.Join(stringx.Remove(pPostContentFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"

	cachePaopaoPPostContentIdPrefix = "cache:paopao:pPostContent:id:"
)

type (
	pPostContentModel interface {
		Insert(ctx context.Context, data *PPostContent) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PPostContent, error)
		Update(ctx context.Context, data *PPostContent) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPPostContentModel struct {
		sqlc.CachedConn
		table string
	}

	PPostContent struct {
		Id         int64  `db:"id"`          // 内容ID
		PostId     int64  `db:"post_id"`     // POST ID
		UserId     int64  `db:"user_id"`     // 用户ID
		Content    string `db:"content"`     // 内容
		Type       int64  `db:"type"`        // 类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源，8收费资源
		Sort       int64  `db:"sort"`        // 排序，越小越靠前
		CreatedOn  int64  `db:"created_on"`  // 创建时间
		ModifiedOn int64  `db:"modified_on"` // 修改时间
		DeletedOn  int64  `db:"deleted_on"`  // 删除时间
		IsDel      int64  `db:"is_del"`      // 是否删除 0 为未删除、1 为已删除
	}
)

func newPPostContentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultPPostContentModel {
	return &defaultPPostContentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`p_post_content`",
	}
}

func (m *defaultPPostContentModel) Delete(ctx context.Context, id int64) error {
	paopaoPPostContentIdKey := fmt.Sprintf("%s%v", cachePaopaoPPostContentIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, paopaoPPostContentIdKey)
	return err
}

func (m *defaultPPostContentModel) FindOne(ctx context.Context, id int64) (*PPostContent, error) {
	paopaoPPostContentIdKey := fmt.Sprintf("%s%v", cachePaopaoPPostContentIdPrefix, id)
	var resp PPostContent
	err := m.QueryRowCtx(ctx, &resp, paopaoPPostContentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pPostContentRows, m.table)
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

func (m *defaultPPostContentModel) Insert(ctx context.Context, data *PPostContent) (sql.Result, error) {
	paopaoPPostContentIdKey := fmt.Sprintf("%s%v", cachePaopaoPPostContentIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pPostContentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.PostId, data.UserId, data.Content, data.Type, data.Sort, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel)
	}, paopaoPPostContentIdKey)
	return ret, err
}

func (m *defaultPPostContentModel) Update(ctx context.Context, data *PPostContent) error {
	paopaoPPostContentIdKey := fmt.Sprintf("%s%v", cachePaopaoPPostContentIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pPostContentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.PostId, data.UserId, data.Content, data.Type, data.Sort, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel, data.Id)
	}, paopaoPPostContentIdKey)
	return err
}

func (m *defaultPPostContentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePaopaoPPostContentIdPrefix, primary)
}

func (m *defaultPPostContentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pPostContentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPPostContentModel) tableName() string {
	return m.table
}
