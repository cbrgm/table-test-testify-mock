# Table Test Testify Mock

Very basic example how to use Go-style table-driven tests with the famous [stretch/testify](https://github.com/stretchr/testify) toolkit with common assertions and mocks that plays nicely with the standard library.

Project Structure:

```bash
.
├── README.md
├── engine
│   ├── engine.go
│   └── engine_mock.go
├── main.go
├── navigation
│   ├── navigation.go
│   └── navigation_mock.go
└── vehicle
    ├── vehicle.go
    └── vehicle_test.go
```

A vehicle has to components `navigation` and `engine`. Both dependencies must be mocked in order to test the `vehicles` `VehicleDiagnostics()` function. Please refer to `vehicle/vehicle_test.go` to see the table-driven tests in action by running `go test -v ./...`.
