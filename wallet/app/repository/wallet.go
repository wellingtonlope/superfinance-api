package repository

import (
	"errors"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
)

var (
	ErrWalletNotFound = errors.New("wallet not found")
)

type Wallet interface {
	New(domain.Wallet) (domain.Wallet, error)
	Update(domain.Wallet) (domain.Wallet, error)
	GetByID(id string) (domain.Wallet, error)
}
