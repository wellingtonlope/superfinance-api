package usecase

import "github.com/wellingtonlope/superfinance-api/wallet/domain"

type WalletOutput struct {
	ID          string
	Name        string
	Description string
}

func NewWalletOutput(wallet domain.Wallet) WalletOutput {
	return WalletOutput{wallet.ID, wallet.Name, wallet.Description}
}
