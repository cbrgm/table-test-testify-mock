package navigation

import "github.com/stretchr/testify/mock"

// go:generate mockery --dir navigation --name Navigation --output navigation/mocks --outpkg mocks --with-expecter
type MockedNavigation struct {
	mock.Mock
}

func (m *MockedNavigation) IsOffline() bool {
	args := m.Called()
	return args.Get(0).(bool)
}
