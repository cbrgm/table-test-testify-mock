package vehicle

import (
	"testing"

	"github.com/cbrgm/table-test-testify-mock/engine"
	"github.com/cbrgm/table-test-testify-mock/navigation"

	genEngineMocks "github.com/cbrgm/table-test-testify-mock/engine/mocks"
	genNavigationMocks "github.com/cbrgm/table-test-testify-mock/navigation/mocks"
)

// TestVehicleDiagnostics uses handwritten mocks for the vehicles subcomponents
func TestVehicleDiagnostics(t *testing.T) {
	type mockedFields struct {
		navigation *navigation.MockedNavigation
		engine     *engine.MockedEngine
	}

	type args struct {
		lowBatteryThreshold int
		isOffline           bool
	}

	testCases := []struct {
		// test case description
		desc string
		// function parameters
		args *args
		// expected function return values
		// in this case error
		result error

		// initializations and assertions for mocked dependencies
		on     func(fields *mockedFields)
		assert func(t *testing.T, fields *mockedFields)
	}{
		{
			desc: "must pass battery level same as threshold",
			args: &args{
				lowBatteryThreshold: 10,    // set the threshold
				isOffline:           false, // setIsOffline
			},
			result: nil,
			on: func(fields *mockedFields) {
				fields.engine.On("ChargeLevel").Return(10)      // mock Chargelevel
				fields.navigation.On("IsOffline").Return(false) // mock IsOffline
			},
			assert: func(t *testing.T, fields *mockedFields) {
				fields.engine.AssertNumberOfCalls(t, "ChargeLevel", 1)
				fields.navigation.AssertNumberOfCalls(t, "IsOffline", 1)
			},
		},
		{
			desc: "must fail with error battery level below threshold",
			args: &args{
				lowBatteryThreshold: 10,    // set the threshold
				isOffline:           false, // setIsOffline
			},
			result: errLowBattery, // we expect an error
			on: func(fields *mockedFields) {
				fields.engine.On("ChargeLevel").Return(1)       // mock Chargelevel
				fields.navigation.On("IsOffline").Return(false) // mock IsOffline
			},
			assert: func(t *testing.T, fields *mockedFields) {
				fields.engine.AssertNumberOfCalls(t, "ChargeLevel", 1)
				fields.navigation.AssertNumberOfCalls(t, "IsOffline", 0)
			},
		},
		{
			desc: "must pass battery level above threshold",
			args: &args{
				lowBatteryThreshold: 10,    // set the threshold
				isOffline:           false, // setIsOffline
			},
			result: nil,
			on: func(fields *mockedFields) {
				fields.engine.On("ChargeLevel").Return(55)
				fields.navigation.On("IsOffline").Return(false)
			},
			assert: func(t *testing.T, fields *mockedFields) {
				fields.engine.AssertNumberOfCalls(t, "ChargeLevel", 1)
				fields.navigation.AssertNumberOfCalls(t, "IsOffline", 1)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mocks := &mockedFields{
				&navigation.MockedNavigation{},
				&engine.MockedEngine{},
			}
			car := NewElectricVehicle(mocks.engine, mocks.navigation)
			if tC.on != nil {
				tC.on(mocks)
			}
			err := car.VehicleDiagnostics(tC.args.lowBatteryThreshold, tC.args.isOffline)
			if err != tC.result {
				t.Errorf("got %v, want %v", err, tC.result)
			}
			if tC.assert != nil {
				tC.assert(t, mocks)
			}
		})
	}
}

// TestVehicleDiagnosticsGenerated uses generated mocks for the vehicles subcomponents
func TestVehicleDiagnosticsGenerated(t *testing.T) {
	type mockedFields struct {
		navigation *genNavigationMocks.Navigation
		engine     *genEngineMocks.Engine
	}

	type args struct {
		lowBatteryThreshold int
		isOffline           bool
	}

	testCases := []struct {
		// test case description
		desc string
		// function parameters
		args *args
		// expected function return values
		// in this case error
		result error

		// initializations and assertions for mocked dependencies
		on     func(fields *mockedFields)
		assert func(t *testing.T, fields *mockedFields)
	}{
		{
			desc: "must pass battery level same as threshold",
			args: &args{
				lowBatteryThreshold: 10,    // set the threshold
				isOffline:           false, // setIsOffline
			},
			result: nil,
			on: func(fields *mockedFields) {
				fields.engine.EXPECT().ChargeLevel().Times(1).Return(10)      // mock Chargelevel
				fields.navigation.EXPECT().IsOffline().Times(1).Return(false) // mock IsOffline
			},
			assert: func(t *testing.T, fields *mockedFields) {
				fields.engine.AssertNumberOfCalls(t, "ChargeLevel", 1)
				fields.navigation.AssertNumberOfCalls(t, "IsOffline", 1)
			},
		},
		{
			desc: "must fail with error battery level below threshold",
			args: &args{
				lowBatteryThreshold: 10,    // set the threshold
				isOffline:           false, // setIsOffline
			},
			result: errLowBattery, // we expect an error
			on: func(fields *mockedFields) {
				fields.engine.EXPECT().ChargeLevel().Times(1).Return(1)       // mock Chargelevel
				fields.navigation.EXPECT().IsOffline().Times(1).Return(false) // mock IsOffline
			},
			assert: func(t *testing.T, fields *mockedFields) {
				fields.engine.AssertNumberOfCalls(t, "ChargeLevel", 1)
				fields.navigation.AssertNumberOfCalls(t, "IsOffline", 0)
			},
		},
		{
			desc: "must pass battery level above threshold",
			args: &args{
				lowBatteryThreshold: 10,    // set the threshold
				isOffline:           false, // setIsOffline
			},
			result: nil,
			on: func(fields *mockedFields) {
				fields.engine.EXPECT().ChargeLevel().Times(1).Return(55)      // mock Chargelevel
				fields.navigation.EXPECT().IsOffline().Times(1).Return(false) // mock IsOffline
			},
			assert: func(t *testing.T, fields *mockedFields) {
				fields.engine.AssertNumberOfCalls(t, "ChargeLevel", 1)
				fields.navigation.AssertNumberOfCalls(t, "IsOffline", 1)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mocks := &mockedFields{
				&genNavigationMocks.Navigation{},
				&genEngineMocks.Engine{},
			}
			car := NewElectricVehicle(mocks.engine, mocks.navigation)
			if tC.on != nil {
				tC.on(mocks)
			}
			err := car.VehicleDiagnostics(tC.args.lowBatteryThreshold, tC.args.isOffline)
			if err != tC.result {
				t.Errorf("got %v, want %v", err, tC.result)
			}
			if tC.assert != nil {
				tC.assert(t, mocks)
			}
		})
	}
}
