package usecase

import "github.com/wellingtonlope/superfinance-api/wallet/app/repository"

type UpdateWalletInput struct {
	ID          string
	Name        string
	Description string
}

type UpdateWallet interface {
	Handle(input UpdateWalletInput) (WalletOutput, error)
}

type updateWallet struct {
	walletRepository repository.Wallet
}

func NewUpdateWallet(walletRepository repository.Wallet) UpdateWallet {
	return &updateWallet{walletRepository}
}

func (uc *updateWallet) Handle(input UpdateWalletInput) (WalletOutput, error) {
	wallet, err := uc.walletRepository.GetByID(input.ID)
	if err != nil {
		return WalletOutput{}, err
	}

	err = wallet.Update(input.Name, input.Description)
	if err != nil {
		return WalletOutput{}, err
	}

	_, err = uc.walletRepository.Update(wallet)
	if err != nil {
		return WalletOutput{}, err
	}

	return NewWalletOutput(wallet), nil
}
