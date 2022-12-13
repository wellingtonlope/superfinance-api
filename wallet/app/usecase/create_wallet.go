package usecase

import (
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
)

type CreateWalletInput struct {
	Name        string
	Description string
}

type CreateWallet interface {
	Handle(CreateWalletInput) (WalletOutput, error)
}

type createWallet struct {
	walletRepository repository.Wallet
}

func NewCreateWallet(walletRepository repository.Wallet) CreateWallet {
	return &createWallet{walletRepository}
}

func (uc *createWallet) Handle(input CreateWalletInput) (WalletOutput, error) {
	wallet, err := domain.NewWallet(input.Name, input.Description)
	if err != nil {
		return WalletOutput{}, err
	}

	walletInserted, err := uc.walletRepository.New(wallet)
	if err != nil {
		return WalletOutput{}, err
	}

	return WalletOutput{walletInserted.ID, wallet.Name, wallet.Description}, nil
}
