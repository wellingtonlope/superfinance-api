package usecase

import (
	"github.com/stretchr/testify/mock"
)

type CreateWalletMock struct {
	mock.Mock
}

func (m *CreateWalletMock) Handle(input CreateWalletInput) (WalletOutput, error) {
	args := m.Called(input)
	var result WalletOutput
	if args.Get(0) != nil {
		result = args.Get(0).(WalletOutput)
	}
	return result, args.Error(1)
}

func NewCreateWalletMock() *CreateWalletMock {
	return new(CreateWalletMock)
}
