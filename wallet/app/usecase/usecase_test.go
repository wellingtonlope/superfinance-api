package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/superfinance-api/wallet/app/usecase"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
	"testing"
)

func TestNewWalletOutput(t *testing.T) {
	expected := usecase.WalletOutput{ID: "1", Name: "wallet name", Description: "wallet description"}
	wallet := domain.Wallet{ID: expected.ID, Name: expected.Name, Description: expected.Description}

	got := usecase.NewWalletOutput(wallet)

	assert.Equal(t, expected, got)
}
