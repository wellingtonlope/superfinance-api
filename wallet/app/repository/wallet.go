package repository

import "github.com/wellingtonlope/superfinance-api/wallet/domain"

type Wallet interface {
	New(domain.Wallet) (domain.Wallet, error)
	Update(domain.Wallet) (domain.Wallet, error)
}
