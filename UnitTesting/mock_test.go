package mockingaround

import (
	"testing"
)

// we are making a mock Something struct with an attribute of mockIncrement which
// has the same function signature as real Something.Increment
// Something Interface is implemented just like the real something
type MockSomething struct {
	mockIncrement func() error
}

//Increment will check to see if we have a defined mockIncrement,
// otherwise return success happy path
func (s *MockSomething) Increment() error {
	if s.mockIncrement != nil {
		return s.mockIncrement()
	}
	return nil
}

func TestActionOnSomething(t *testing.T) {
	// create a MockSomething, we will use as a drop in replacement
	// of Something in our function call we are trying to test
	ms := &MockSomething{
	// mockIncrement: func() error {
	// 	// this time we want to mock our a failure scenario
	// 	return errors.New("failed to increment")
	// },
	}
	// perform the function call we want to test
	actionOnSomething(ms)
	// ... check for panic
}
