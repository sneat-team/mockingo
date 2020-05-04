package mockingo

import (
	"fmt"
)

type Argument struct {
	Name  string
	Value interface{}
}

type MethodCall struct {
	Args map[string]interface{}
}

type MockedMethod struct {
	name  string
	calls []MethodCall
}

func (v *MockedMethod) Name() string {
	return v.name
}

func (v *MockedMethod) Called(args ...Argument) {
	var methodCall = MethodCall{Args: make(map[string]interface{}, len(args))}
	for _, arg := range args {
		methodCall.Args[arg.Name] = arg.Value
	}
	v.calls = append(v.calls, methodCall)
}

func (v *MockedMethod) Calls() []MethodCall {
	return v.calls[:]
}

func NewMockedMethod(name string) *MockedMethod {
	return &MockedMethod{name: name}
}

func (v *MockedMethod) AssertCalledExactly(expected int) string {
	if n := len(v.calls); n != expected {
		return fmt.Sprintf("expected to get %v calls to %v(), got: %v", expected, v.name, n)
	}
	return ""
}

func (v *MockedMethod) AssertCalledAtLeastOnce() string {
	if n := len(v.calls); n == 0 {
		return fmt.Sprintf("expected to get at least 1 call to %v(), got none", v.name)
	}
	return ""
}
