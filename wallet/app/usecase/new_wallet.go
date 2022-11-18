package usecase

import (
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
)

type NewWalletInput struct {
	Name        string
	Description string
}

type NewWalletOutput struct {
	ID          string
	Name        string
	Description string
}

func NewNewWalletOutput(wallet domain.Wallet) NewWalletOutput {
	return NewWalletOutput{wallet.ID, wallet.Name, wallet.Description}
}

type NewWallet interface {
	Handle(NewWalletInput) (NewWalletOutput, error)
}

type newWallet struct {
	walletRepository repository.Wallet
}

func NewNewWallet(walletRepository repository.Wallet) NewWallet {
	return &newWallet{walletRepository}
}

func (uc *newWallet) Handle(input NewWalletInput) (NewWalletOutput, error) {
	wallet, err := domain.NewWallet(input.Name, input.Description)
	if err != nil {
		return NewWalletOutput{}, err
	}

	walletInserted, err := uc.walletRepository.New(wallet)
	if err != nil {
		return NewWalletOutput{}, err
	}

	return NewWalletOutput{walletInserted.ID, wallet.Name, wallet.Description}, nil
}
