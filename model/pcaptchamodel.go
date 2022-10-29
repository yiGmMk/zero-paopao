package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PCaptchaModel = (*customPCaptchaModel)(nil)

type (
	// PCaptchaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPCaptchaModel.
	PCaptchaModel interface {
		pCaptchaModel
	}

	customPCaptchaModel struct {
		*defaultPCaptchaModel
	}
)

// NewPCaptchaModel returns a model for the database table.
func NewPCaptchaModel(conn sqlx.SqlConn) PCaptchaModel {
	return &customPCaptchaModel{
		defaultPCaptchaModel: newPCaptchaModel(conn),
	}
}
