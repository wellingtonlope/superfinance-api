package usecase

import "github.com/stretchr/testify/mock"

type GetByIDWalletMock struct {
	mock.Mock
}

func NewGetByIDWalletMock() *GetByIDWalletMock {
	return new(GetByIDWalletMock)
}

func (m *GetByIDWalletMock) Handle(id string) (WalletOutput, error) {
	args := m.Called(id)
	var result WalletOutput
	if args.Get(0) != nil {
		result = args.Get(0).(WalletOutput)
	}
	return result, args.Error(1)
}
