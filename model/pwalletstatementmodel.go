package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PWalletStatementModel = (*customPWalletStatementModel)(nil)

type (
	// PWalletStatementModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPWalletStatementModel.
	PWalletStatementModel interface {
		pWalletStatementModel
	}

	customPWalletStatementModel struct {
		*defaultPWalletStatementModel
	}
)

// NewPWalletStatementModel returns a model for the database table.
func NewPWalletStatementModel(conn sqlx.SqlConn) PWalletStatementModel {
	return &customPWalletStatementModel{
		defaultPWalletStatementModel: newPWalletStatementModel(conn),
	}
}
