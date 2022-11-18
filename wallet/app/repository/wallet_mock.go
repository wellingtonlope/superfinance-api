package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
)

type WalletMock struct {
	mock.Mock
}

func (m *WalletMock) New(wallet domain.Wallet) (domain.Wallet, error) {
	args := m.Called(wallet)
	var result domain.Wallet
	if args.Get(0) != nil {
		result = args.Get(0).(domain.Wallet)
	}
	return result, args.Error(1)
}

func (m *WalletMock) Update(wallet domain.Wallet) (domain.Wallet, error) {
	args := m.Called(wallet)
	var result domain.Wallet
	if args.Get(0) != nil {
		result = args.Get(0).(domain.Wallet)
	}
	return result, args.Error(1)
}
