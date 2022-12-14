package usecase

import "github.com/stretchr/testify/mock"

type UpdateWalletMock struct {
	mock.Mock
}

func NewUpdateWalletMock() *UpdateWalletMock {
	return new(UpdateWalletMock)
}

func (m *UpdateWalletMock) Handle(input UpdateWalletInput) (WalletOutput, error) {
	args := m.Called(input)
	var result WalletOutput
	if args.Get(0) != nil {
		result = args.Get(0).(WalletOutput)
	}
	return result, args.Error(1)
}
