package mockingo

import (
	"fmt"
)

type MockedFunc struct {
	name  string
	calls []MethodCall
}

func (v *MockedFunc) Name() string {
	return v.name
}

func (v *MockedFunc) Called(args ...Argument) {
	var methodCall = MethodCall{args: make(map[string]interface{}, len(args))}
	for _, arg := range args {
		methodCall.args[arg.name] = arg.Value
	}
	v.calls = append(v.calls, methodCall)
}

func (v *MockedFunc) Calls() []MethodCall {
	calls := make([]MethodCall, len(v.calls))
	copy(calls, v.calls)
	return calls
}

func NewMockedMethod(name string) *MockedFunc {
	return &MockedFunc{name: name}
}

func (v *MockedFunc) AssertCalledExactly(expected int) string {
	if n := len(v.calls); n != expected {
		return fmt.Sprintf("expected to get %v calls to %v(), got: %v", expected, v.name, n)
	}
	return ""
}

func (v *MockedFunc) AssertCalledAtLeastOnce() string {
	if n := len(v.calls); n == 0 {
		return fmt.Sprintf("expected to get at least 1 call to %v(), got none", v.name)
	}
	return ""
}
