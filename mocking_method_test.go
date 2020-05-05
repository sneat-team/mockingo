package mockingo

import "testing"

func TestNewmockedFunc(t *testing.T) {
	mockedFunc := NewMockedFunc(t, "test1")
	if mockedFunc == nil {
		t.Fatal("mockedFunc == nil")
	}
	if len(mockedFunc.Calls()) != 0 {
		t.Fatal("len(mockedFunc.Calls()) != 0")
	}
	mockedFunc.Called()
	if len(mockedFunc.Calls()) != 1 {
		t.Fatal("len(mockedFunc.Calls()) != 1")
	}
}

func TestMockedFunc_AssertCalledExactly(t *testing.T) {
	mockedFunc := NewMockedFunc(t, "test1")
	mockedFunc.Called()
	mockedFunc.AssertCalledExactly(1)
}

func TestMockedFunc_AssertCalledAtLeastOnce(t *testing.T) {
	mockedFunc := NewMockedFunc(t, "test1")
	mockedFunc.Called()
	mockedFunc.AssertCalledAtLeastOnce()
}

func TestMockedFunc_Calls(t *testing.T) {
	mockedFunc := NewMockedFunc(t, "test1")
	mockedFunc.Called()
	if n := len(mockedFunc.calls); n != 1 {
		t.Fatalf("expected to have 1 call, got: %v", n)
	}
	if n := len(mockedFunc.calls[0].args); n > 0 {
		t.Fatalf("expected to have no arguments stored for the 1st call, got %v", n)
	}

	mockedFunc.Called(
		NewArgument("a1", 11),
		NewArgument("a2", 22),
	)
	
	if n := len(mockedFunc.calls); n != 2 {
		t.Errorf("expected to have 2 calls, got: %v", n)
	}
	if n := len(mockedFunc.calls[1].args); n != 2 {
		t.Fatalf("expected to have 2 arguments stored for the 2nd call, got %v", n)
	}
	for i, a := range []string{"a1", "a2"} {
		if v, ok := mockedFunc.calls[1].args[a]; !ok {
			t.Errorf("expected argument '%v' not found", a)
		} else {
			if expected := (i+1)*10 + i + 1; v.(int) != expected {
				t.Errorf("argument '%v' expected to keep value %v but got %v", a, expected, v)
			}
		}
	}
}
