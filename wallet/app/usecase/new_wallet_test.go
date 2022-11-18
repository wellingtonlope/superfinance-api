package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/app/usecase"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
)

func TestNewNewWalletOutput(t *testing.T) {
	expected := usecase.NewWalletOutput{"1", "wallet name", "wallet description"}
	wallet := domain.Wallet{ID: expected.ID, Name: expected.Name, Description: expected.Description}

	got := usecase.NewNewWalletOutput(wallet)

	assert.Equal(t, expected, got)
}

func TestNewWallet(t *testing.T) {
	walletRepository := new(repository.WalletMock)
	uc := usecase.NewNewWallet(walletRepository)

	input := usecase.NewWalletInput{"wallet name", "wallet description"}
	walletRepository.
		On("New", domain.Wallet{Name: input.Name, Description: input.Description}).
		Return(domain.Wallet{ID: "1", Name: input.Name, Description: input.Description}, nil)

	got, err := uc.Handle(input)

	assert.Nil(t, err)
	assert.NotEmpty(t, got.ID)
	assert.Equal(t, input.Name, got.Name)
	assert.Equal(t, input.Description, got.Description)
}
