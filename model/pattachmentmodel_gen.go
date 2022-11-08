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
	pAttachmentFieldNames          = builder.RawFieldNames(&PAttachment{})
	pAttachmentRows                = strings.Join(pAttachmentFieldNames, ",")
	pAttachmentRowsExpectAutoSet   = strings.Join(stringx.Remove(pAttachmentFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	pAttachmentRowsWithPlaceHolder = strings.Join(stringx.Remove(pAttachmentFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"

	cachePaopaoPAttachmentIdPrefix = "cache:paopao:pAttachment:id:"
)

type (
	pAttachmentModel interface {
		Insert(ctx context.Context, data *PAttachment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PAttachment, error)
		Update(ctx context.Context, data *PAttachment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPAttachmentModel struct {
		sqlc.CachedConn
		table string
	}

	PAttachment struct {
		Id         int64  `db:"id"`
		UserId     int64  `db:"user_id"`
		FileSize   int64  `db:"file_size"`
		ImgWidth   int64  `db:"img_width"`
		ImgHeight  int64  `db:"img_height"`
		Type       int64  `db:"type"` // 1图片，2视频，3其他附件
		Content    string `db:"content"`
		CreatedOn  int64  `db:"created_on"`  // 创建时间
		ModifiedOn int64  `db:"modified_on"` // 修改时间
		DeletedOn  int64  `db:"deleted_on"`  // 删除时间
		IsDel      int64  `db:"is_del"`      // 是否删除 0 为未删除、1 为已删除
	}
)

func newPAttachmentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultPAttachmentModel {
	return &defaultPAttachmentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`p_attachment`",
	}
}

func (m *defaultPAttachmentModel) Delete(ctx context.Context, id int64) error {
	paopaoPAttachmentIdKey := fmt.Sprintf("%s%v", cachePaopaoPAttachmentIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, paopaoPAttachmentIdKey)
	return err
}

func (m *defaultPAttachmentModel) FindOne(ctx context.Context, id int64) (*PAttachment, error) {
	paopaoPAttachmentIdKey := fmt.Sprintf("%s%v", cachePaopaoPAttachmentIdPrefix, id)
	var resp PAttachment
	err := m.QueryRowCtx(ctx, &resp, paopaoPAttachmentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pAttachmentRows, m.table)
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

func (m *defaultPAttachmentModel) Insert(ctx context.Context, data *PAttachment) (sql.Result, error) {
	paopaoPAttachmentIdKey := fmt.Sprintf("%s%v", cachePaopaoPAttachmentIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pAttachmentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.FileSize, data.ImgWidth, data.ImgHeight, data.Type, data.Content, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel)
	}, paopaoPAttachmentIdKey)
	return ret, err
}

func (m *defaultPAttachmentModel) Update(ctx context.Context, data *PAttachment) error {
	paopaoPAttachmentIdKey := fmt.Sprintf("%s%v", cachePaopaoPAttachmentIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pAttachmentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.FileSize, data.ImgWidth, data.ImgHeight, data.Type, data.Content, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel, data.Id)
	}, paopaoPAttachmentIdKey)
	return err
}

func (m *defaultPAttachmentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePaopaoPAttachmentIdPrefix, primary)
}

func (m *defaultPAttachmentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pAttachmentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPAttachmentModel) tableName() string {
	return m.table
}
