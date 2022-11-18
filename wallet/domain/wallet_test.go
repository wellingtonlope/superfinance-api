package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
)

func TestNewWallet(t *testing.T) {
	testsCase := []struct {
		name                string
		expectedName        string
		expectedDescription string
		expectedError       error
	}{
		{
			"should create a new wallet",
			"wallet name",
			"wallet description",
			nil,
		},
		{
			"should create a new wallet when description is empty",
			"wallet name",
			"",
			nil,
		},
		{
			"shouldn't create a new wallet when name is empty",
			"",
			"wallet description",
			domain.ErrWalletNameEmpty,
		},
	}

	for _, test := range testsCase {
		t.Run(test.name, func(t *testing.T) {
			got, err := domain.NewWallet(test.expectedName, test.expectedDescription)

			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
				assert.Equal(t, domain.Wallet{}, got)
				return
			}

			assert.Empty(t, got.ID)
			assert.Equal(t, test.expectedName, got.Name)
			assert.Equal(t, test.expectedDescription, got.Description)
		})
	}
}

func TestUpdateWallet(t *testing.T) {
	testsCase := []struct {
		name                string
		expectedName        string
		expectedDescription string
		expectedError       error
		wallet              domain.Wallet
	}{
		{
			"should update an wallet",
			"wallet name updated",
			"wallet description updated",
			nil,
			domain.Wallet{"1", "wallet name", "wallet description"},
		},
		{
			"should update an wallet when description is empty",
			"wallet name updated",
			"",
			nil,
			domain.Wallet{"1", "wallet name", "wallet description"},
		},
		{
			"shouldn't update an wallet when name is empty",
			"",
			"wallet description updated",
			domain.ErrWalletNameEmpty,
			domain.Wallet{"1", "wallet name", "wallet description"},
		},
	}

	for _, test := range testsCase {
		t.Run(test.name, func(t *testing.T) {
			walletBeforeUpdate := test.wallet
			err := test.wallet.UpdateWallet(test.expectedName, test.expectedDescription)

			if err != nil {
				assert.Equal(t, domain.ErrWalletNameEmpty, err)
				assert.Equal(t, walletBeforeUpdate, test.wallet)
				return
			}

			assert.Equal(t, walletBeforeUpdate.ID, test.wallet.ID)
			assert.Equal(t, test.expectedName, test.wallet.Name)
			assert.Equal(t, test.expectedDescription, test.wallet.Description)
		})
	}
}
