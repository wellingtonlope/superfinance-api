package memory_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/superfinance-api/wallet/app/repository"
	"github.com/wellingtonlope/superfinance-api/wallet/domain"
	"github.com/wellingtonlope/superfinance-api/wallet/infra/repository/memory"
	"testing"
)

func TestWallet_New(t *testing.T) {
	rep := memory.NewWallet()

	input := domain.Wallet{Name: "name", Description: "description"}
	got, err := rep.New(input)

	assert.Nil(t, err)
	assert.NotEmpty(t, got.ID)
	assert.Equal(t, input.Name, got.Name)
	assert.Equal(t, input.Description, got.Description)
}

func TestWallet_GetByID(t *testing.T) {
	t.Run("should get wallet by ID", func(t *testing.T) {
		rep := memory.NewWallet()
		inserted, _ := rep.New(domain.Wallet{Name: "name", Description: "description"})

		got, err := rep.GetByID(inserted.ID)

		assert.Nil(t, err)
		assert.Equal(t, inserted, got)
	})

	t.Run("should return not found when wallet not exists", func(t *testing.T) {
		rep := memory.NewWallet()

		got, err := rep.GetByID("no-exists")

		assert.Empty(t, got)
		assert.Equal(t, repository.ErrWalletNotFound, err)
	})
}

func TestWallet_Update(t *testing.T) {
	t.Run("should update the wallet by ID", func(t *testing.T) {
		rep := memory.NewWallet()
		inserted, _ := rep.New(domain.Wallet{Name: "name", Description: "description"})

		inputUpdate := domain.Wallet{ID: inserted.ID, Name: "name updated", Description: "description updated"}
		updated, err := rep.Update(inputUpdate)

		updatedCheck, _ := rep.GetByID(updated.ID)
		assert.Nil(t, err)
		assert.Equal(t, inputUpdate, updated)
		assert.Equal(t, inputUpdate, updatedCheck)
	})

	t.Run("should update the wallet that not exists", func(t *testing.T) {
		rep := memory.NewWallet()

		inputUpdate := domain.Wallet{ID: "1", Name: "name updated", Description: "description updated"}
		updated, err := rep.Update(inputUpdate)

		assert.Equal(t, repository.ErrWalletNotFound, err)
		assert.Empty(t, updated)
	})
}
