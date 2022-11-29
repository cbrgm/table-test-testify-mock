package navigation

import "github.com/stretchr/testify/mock"

type MockedNavigation struct {
	mock.Mock
}

func (m *MockedNavigation) IsOffline() bool {
	args := m.Called()
	return args.Get(0).(bool)
}
