package domain

import (
	"errors"
)

var (
	ErrWalletNameEmpty = errors.New("wallet name mustn't be empty")
)

type Wallet struct {
	ID          string
	Name        string
	Description string
}

func NewWallet(name, description string) (Wallet, error) {
	if name == "" {
		return Wallet{}, ErrWalletNameEmpty
	}
	return Wallet{ID: "", Name: name, Description: description}, nil
}

func (w *Wallet) Update(name, description string) error {
	if name == "" {
		return ErrWalletNameEmpty
	}

	w.Name = name
	w.Description = description
	return nil
}
