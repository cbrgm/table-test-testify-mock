package navigation

type Navigation interface {
	IsOffline() bool // whether to use offline navigation or not
}

type HereMaps struct {
	Offline bool
}

func (m *HereMaps) IsOffline() bool {
	return m.Offline
}
