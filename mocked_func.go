package mockingo

import (
	"testing"
)

// MockedFunc stores information about call to a mocked func and provide helpers to assert number of calls and arguments
type MockedFunc struct {
	t     *testing.T
	name  string
	calls []MethodCall
}

// Name of a mocked func used to log in error messages
func (v *MockedFunc) Name() string {
	return v.name
}

// Called stores information about a call to a mocked func
func (v *MockedFunc) Called(args ...Argument) {
	var methodCall = MethodCall{args: make(map[string]interface{}, len(args))}
	for _, arg := range args {
		methodCall.args[arg.name] = arg.value
	}
	v.calls = append(v.calls, methodCall)
}

// Returns all calls made to a mocked func
func (v *MockedFunc) Calls() []MethodCall {
	calls := make([]MethodCall, len(v.calls))
	copy(calls, v.calls)
	return calls
}

// NewMockedFunc creates new mocked func store/counters
func NewMockedFunc(t *testing.T, name string) *MockedFunc {
	return &MockedFunc{t: t, name: name}
}

// AssertCalledExactly verifies exactly specified number of call have been made to the mocked func
func (v *MockedFunc) AssertCalledExactly(expected int, fatal ...bool) {
	if n := len(v.calls); n != expected {
		v.t.Helper()
		const m = "expected to get %v calls to %v(), got: %v"
		if len(fatal) > 0 && fatal[0] {
			v.t.Fatalf(m, expected, v.name, n)
		} else {
			v.t.Errorf(m, expected, v.name, n)
		}
	}
}

// AssertCalledExactly verifies at least 1 call have been made to the mocked func
func (v *MockedFunc) AssertCalledAtLeastOnce(fatal ...bool) {
	if n := len(v.calls); n == 0 {
		v.t.Helper()
		const m = "expected to get at least 1 call to %v(), got none"
		if len(fatal) > 0 && fatal[0] {
			v.t.Fatalf(m, v.name)
		} else {
			v.t.Errorf(m, v.name)
		}
	}
}
