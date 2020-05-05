package mockingo

import (
	"testing"
)

type MockedFunc struct {
	t     *testing.T
	name  string
	calls []MethodCall
}

func (v *MockedFunc) Name() string {
	return v.name
}

func (v *MockedFunc) Called(args ...Argument) {
	var methodCall = MethodCall{args: make(map[string]interface{}, len(args))}
	for _, arg := range args {
		methodCall.args[arg.name] = arg.value
	}
	v.calls = append(v.calls, methodCall)
}

func (v *MockedFunc) Calls() []MethodCall {
	calls := make([]MethodCall, len(v.calls))
	copy(calls, v.calls)
	return calls
}

func NewMockedFunc(t *testing.T, name string) *MockedFunc {
	return &MockedFunc{t: t, name: name}
}

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
