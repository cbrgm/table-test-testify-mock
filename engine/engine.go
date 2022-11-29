package engine

import (
	"log"
	"math/rand"
	"time"
)

// go:generate mockery --dir engine --name Engine --output engine/mocks --outpkg mocks --with-expecter
type Engine interface {
	ChargeLevel() int
}

type ElectricEngine struct{}

func (e *ElectricEngine) ChargeLevel() int {
	rand.Seed(time.Now().UnixNano())
	utilization := rand.Intn(100)
	log.Printf("Battery charge level: %v", utilization)
	return utilization
}
