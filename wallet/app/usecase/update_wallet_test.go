package usecase_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/app/usecase"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
	"testing"
)

func TestUpdateWallet(t *testing.T) {
	t.Run("should update a wallet", func(t *testing.T) {
		walletRepository := repository.NewWalletMock()
		walletRepository.
			On("GetByID", "1").
			Return(domain.Wallet{ID: "1", Name: "Wallet name", Description: "Wallet description"}, nil).Once()
		walletRepository.
			On("Update", domain.Wallet{ID: "1", Name: "Wallet name updated", Description: "Wallet description updated"}).
			Return(domain.Wallet{ID: "1", Name: "Wallet name updated", Description: "Wallet description updated"}, nil).Once()
		uc := usecase.NewUpdateWallet(walletRepository)

		input := usecase.UpdateWalletInput{ID: "1", Name: "Wallet name updated", Description: "Wallet description updated"}
		got, err := uc.Handle(input)

		expected := usecase.WalletOutput{ID: "1", Name: "Wallet name updated", Description: "Wallet description updated"}
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
		walletRepository.AssertExpectations(t)
	})

	t.Run("should fail when get by id return error", func(t *testing.T) {
		walletRepository := repository.NewWalletMock()
		expectedErr := repository.ErrWalletNotFound
		walletRepository.
			On("GetByID", "1").
			Return(domain.Wallet{}, expectedErr).Once()
		uc := usecase.NewUpdateWallet(walletRepository)

		input := usecase.UpdateWalletInput{ID: "1", Name: "Wallet name updated", Description: "Wallet description updated"}
		got, err := uc.Handle(input)

		assert.Equal(t, expectedErr, err)
		assert.Empty(t, got)
		walletRepository.AssertExpectations(t)
	})

	t.Run("should fail when domain update return error", func(t *testing.T) {
		walletRepository := repository.NewWalletMock()
		walletRepository.
			On("GetByID", "1").
			Return(domain.Wallet{ID: "1", Name: "Wallet name", Description: "Wallet description"}, nil).Once()
		uc := usecase.NewUpdateWallet(walletRepository)

		input := usecase.UpdateWalletInput{ID: "1", Name: "", Description: "Wallet description updated"}
		got, err := uc.Handle(input)

		assert.Equal(t, domain.ErrWalletNameEmpty, err)
		assert.Empty(t, got)
		walletRepository.AssertExpectations(t)
	})

	t.Run("should fail when repository update return error", func(t *testing.T) {
		walletRepository := repository.NewWalletMock()
		expectedErr := errors.New("i'm an error")
		walletRepository.
			On("GetByID", "1").
			Return(domain.Wallet{ID: "1", Name: "Wallet name", Description: "Wallet description"}, nil).Once()
		walletRepository.
			On("Update", domain.Wallet{ID: "1", Name: "Wallet name updated", Description: "Wallet description updated"}).
			Return(domain.Wallet{}, expectedErr).Once()
		uc := usecase.NewUpdateWallet(walletRepository)

		input := usecase.UpdateWalletInput{ID: "1", Name: "Wallet name updated", Description: "Wallet description updated"}
		got, err := uc.Handle(input)

		assert.Equal(t, expectedErr, err)
		assert.Empty(t, got)
		walletRepository.AssertExpectations(t)
	})
}
