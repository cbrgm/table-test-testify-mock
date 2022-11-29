package engine

import (
	"github.com/stretchr/testify/mock"
)

type MockedEngine struct {
	mock.Mock
}

func (m *MockedEngine) ChargeLevel() int {
	args := m.Called()
	return args.Get(0).(int)
}
