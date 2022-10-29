package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PWalletRechargeModel = (*customPWalletRechargeModel)(nil)

type (
	// PWalletRechargeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPWalletRechargeModel.
	PWalletRechargeModel interface {
		pWalletRechargeModel
	}

	customPWalletRechargeModel struct {
		*defaultPWalletRechargeModel
	}
)

// NewPWalletRechargeModel returns a model for the database table.
func NewPWalletRechargeModel(conn sqlx.SqlConn) PWalletRechargeModel {
	return &customPWalletRechargeModel{
		defaultPWalletRechargeModel: newPWalletRechargeModel(conn),
	}
}
