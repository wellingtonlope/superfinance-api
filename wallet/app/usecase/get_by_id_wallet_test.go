package usecase_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/app/usecase"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
	"testing"
)

func TestGetByIDWallet_Handle(t *testing.T) {
	t.Run("should get wallet by ID", func(t *testing.T) {
		walletRepository := repository.NewWalletMock()
		uc := usecase.NewGetByIDWallet(walletRepository)
		expected := usecase.WalletOutput{ID: "1", Name: "name wallet", Description: "description wallet"}
		walletRepository.
			On("GetByID", expected.ID).
			Return(domain.Wallet{ID: expected.ID, Name: expected.Name, Description: expected.Description}, nil)

		got, err := uc.Handle(expected.ID)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("should fail when repository return error", func(t *testing.T) {
		walletRepository := repository.NewWalletMock()
		uc := usecase.NewGetByIDWallet(walletRepository)
		expectedErr := errors.New("i'm an error")
		walletRepository.
			On("GetByID", "1").
			Return(domain.Wallet{}, expectedErr)

		got, err := uc.Handle("1")
		assert.Equal(t, expectedErr, err)
		assert.Empty(t, got)
	})
}
