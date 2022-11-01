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
	pMessageFieldNames          = builder.RawFieldNames(&PMessage{})
	pMessageRows                = strings.Join(pMessageFieldNames, ",")
	pMessageRowsExpectAutoSet   = strings.Join(stringx.Remove(pMessageFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	pMessageRowsWithPlaceHolder = strings.Join(stringx.Remove(pMessageFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
)

type (
	pMessageModel interface {
		Insert(ctx context.Context, data *PMessage) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PMessage, error)
		Update(ctx context.Context, data *PMessage) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPMessageModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PMessage struct {
		Id             int64  `db:"id"`               // 消息通知ID
		SenderUserId   int64  `db:"sender_user_id"`   // 发送方用户ID
		ReceiverUserId int64  `db:"receiver_user_id"` // 接收方用户ID
		Type           int64  `db:"type"`             // 通知类型，1动态，2评论，3回复，4私信，99系统通知
		Brief          string `db:"brief"`            // 摘要说明
		Content        string `db:"content"`          // 详细内容
		PostId         int64  `db:"post_id"`          // 动态ID
		CommentId      int64  `db:"comment_id"`       // 评论ID
		ReplyId        int64  `db:"reply_id"`         // 回复ID
		IsRead         int64  `db:"is_read"`          // 是否已读
		CreatedOn      int64  `db:"created_on"`       // 创建时间
		ModifiedOn     int64  `db:"modified_on"`      // 修改时间
		DeletedOn      int64  `db:"deleted_on"`       // 删除时间
		IsDel          int64  `db:"is_del"`           // 是否删除 0 为未删除、1 为已删除
	}
)

func newPMessageModel(conn sqlx.SqlConn) *defaultPMessageModel {
	return &defaultPMessageModel{
		conn:  conn,
		table: "`p_message`",
	}
}

func (m *defaultPMessageModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPMessageModel) FindOne(ctx context.Context, id int64) (*PMessage, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pMessageRows, m.table)
	var resp PMessage
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

func (m *defaultPMessageModel) Insert(ctx context.Context, data *PMessage) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pMessageRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.SenderUserId, data.ReceiverUserId, data.Type, data.Brief, data.Content, data.PostId, data.CommentId, data.ReplyId, data.IsRead, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel)
	return ret, err
}

func (m *defaultPMessageModel) Update(ctx context.Context, data *PMessage) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pMessageRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.SenderUserId, data.ReceiverUserId, data.Type, data.Brief, data.Content, data.PostId, data.CommentId, data.ReplyId, data.IsRead, data.CreatedOn, data.ModifiedOn, data.DeletedOn, data.IsDel, data.Id)
	return err
}

func (m *defaultPMessageModel) tableName() string {
	return m.table
}