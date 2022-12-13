package usecase

import "github.com/wellingtonlope/superfinance-api/wallet/app/repository"

type GeByIDWallet interface {
	Handle(id string) (WalletOutput, error)
}

type getByIDWallet struct {
	walletRepository repository.Wallet
}

func NewGetByIDWallet(walletRepository repository.Wallet) GeByIDWallet {
	return &getByIDWallet{walletRepository}
}

func (uc *getByIDWallet) Handle(id string) (WalletOutput, error) {
	wallet, err := uc.walletRepository.GetByID(id)
	if err != nil {
		return WalletOutput{}, err
	}
	return NewWalletOutput(wallet), nil
}
