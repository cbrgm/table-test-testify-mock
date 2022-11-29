package vehicle

import (
	"errors"

	"github.com/cbrgm/table-test-testify-mock/engine"
	"github.com/cbrgm/table-test-testify-mock/navigation"
)

var (
	errLowBattery            = errors.New("battery level is critical")
	errNoNetworkConnectivity = errors.New("network has connection issues")
)

type ElectricVehicle struct {
	engine     engine.Engine
	navigation navigation.Navigation
}

func NewElectricVehicle(engine engine.Engine, navigation navigation.Navigation) *ElectricVehicle {
	return &ElectricVehicle{
		engine:     engine,
		navigation: navigation,
	}
}

func (e *ElectricVehicle) VehicleDiagnostics(chargingThreshold int, isOffline bool) error {
	if e.engine.ChargeLevel() < chargingThreshold {
		return errLowBattery
	}

	if e.navigation.IsOffline() != isOffline {
		return errNoNetworkConnectivity
	}

	return nil
}
