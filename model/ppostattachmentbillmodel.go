package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PPostAttachmentBillModel = (*customPPostAttachmentBillModel)(nil)

type (
	// PPostAttachmentBillModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPPostAttachmentBillModel.
	PPostAttachmentBillModel interface {
		pPostAttachmentBillModel
	}

	customPPostAttachmentBillModel struct {
		*defaultPPostAttachmentBillModel
	}
)

// NewPPostAttachmentBillModel returns a model for the database table.
func NewPPostAttachmentBillModel(conn sqlx.SqlConn) PPostAttachmentBillModel {
	return &customPPostAttachmentBillModel{
		defaultPPostAttachmentBillModel: newPPostAttachmentBillModel(conn),
	}
}
