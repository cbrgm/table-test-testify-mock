package main

import (
	"github.com/cbrgm/table-test-testify-mock/engine"
	"github.com/cbrgm/table-test-testify-mock/navigation"
	"github.com/cbrgm/table-test-testify-mock/vehicle"
)

func main() {
	navigationComponent := &navigation.HereMaps{}
	engineComponent := &engine.ElectricEngine{}
	// Constructor receives types implementing required interfaces
	_ = vehicle.NewElectricVehicle(engineComponent, navigationComponent)
}
