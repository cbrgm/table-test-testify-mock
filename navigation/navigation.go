package navigation

// go:generate mockery --dir navigation --name Navigation --output navigation/mocks --outpkg mocks --with-expecter
type Navigation interface {
	IsOffline() bool // whether to use offline navigation or not
}

type HereMaps struct {
	Offline bool
}

func (m *HereMaps) IsOffline() bool {
	return m.Offline
}
