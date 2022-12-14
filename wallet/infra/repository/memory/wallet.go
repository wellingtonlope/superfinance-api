package memory

import (
	"github.com/google/uuid"
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
)

type wallet struct {
	db map[string]domain.Wallet
}

func NewWallet() repository.Wallet {
	return &wallet{
		db: map[string]domain.Wallet{},
	}
}

func (r *wallet) New(wallet domain.Wallet) (domain.Wallet, error) {
	id := uuid.NewString()
	wallet.ID = id
	r.db[id] = wallet
	return wallet, nil
}

func (r *wallet) Update(wallet domain.Wallet) (domain.Wallet, error) {
	_, err := r.GetByID(wallet.ID)
	if err != nil {
		return domain.Wallet{}, err
	}

	r.db[wallet.ID] = wallet

	return wallet, nil
}

func (r *wallet) GetByID(ID string) (domain.Wallet, error) {
	if result, ok := r.db[ID]; ok {
		return result, nil
	}

	return domain.Wallet{}, repository.ErrWalletNotFound
}
