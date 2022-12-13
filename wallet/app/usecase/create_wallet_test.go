package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/app/usecase"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
)

func TestCreateWallet(t *testing.T) {
	t.Run("should create a wallet", func(t *testing.T) {
		walletRepository := repository.NewWalletMock()
		input := usecase.CreateWalletInput{Name: "wallet name", Description: "wallet description"}
		walletRepository.
			On("New", domain.Wallet{Name: input.Name, Description: input.Description}).
			Return(domain.Wallet{ID: "1", Name: input.Name, Description: input.Description}, nil)

		uc := usecase.NewCreateWallet(walletRepository)
		got, err := uc.Handle(input)

		assert.Nil(t, err)
		assert.NotEmpty(t, got.ID)
		assert.Equal(t, input.Name, got.Name)
		assert.Equal(t, input.Description, got.Description)
	})

	t.Run("should fail when domain new return error", func(t *testing.T) {
		walletRepository := repository.NewWalletMock()

		uc := usecase.NewCreateWallet(walletRepository)
		input := usecase.CreateWalletInput{Name: "", Description: "wallet description"}
		got, err := uc.Handle(input)

		assert.Equal(t, domain.ErrWalletNameEmpty, err)
		assert.Empty(t, got)
	})

	t.Run("should fail when repository new return error", func(t *testing.T) {
		walletRepository := repository.NewWalletMock()
		expectedErr := errors.New("i'm an error")
		input := usecase.CreateWalletInput{Name: "wallet name", Description: "wallet description"}
		walletRepository.
			On("New", domain.Wallet{Name: input.Name, Description: input.Description}).
			Return(domain.Wallet{}, expectedErr)

		uc := usecase.NewCreateWallet(walletRepository)
		got, err := uc.Handle(input)

		assert.Equal(t, expectedErr, err)
		assert.Empty(t, got)
	})
}
