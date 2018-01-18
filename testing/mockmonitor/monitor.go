package mockmonitor

import (
	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/mock"
)

// MockMonitor ...
type MockMonitor struct {
	mock.Mock
}

// Stop ...
func (m *MockMonitor) Stop() {
	return
}

// IsStopped ...
func (m *MockMonitor) IsStopped() bool {
	args := m.Called()
	return args.Bool(0)
}

// HandleUpdate ...
func (m *MockMonitor) HandleUpdate(e *types.Event) error {
	args := m.Called(e)
	return args.Error(0)
}

// HandleFailure ...
func (m *MockMonitor) HandleFailure(e *types.Entity) error {
	args := m.Called(e)
	return args.Error(0)
}
