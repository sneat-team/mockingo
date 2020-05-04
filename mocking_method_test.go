package mockingo

import "testing"

func TestNewMockedMethod(t *testing.T) {
	mockedMethod := NewMockedMethod("test1")
	if mockedMethod == nil {
		t.Fatal("mockedMethod == nil")
	}
	if len(mockedMethod.Calls()) != 0 {
		t.Fatal("len(mockedMethod.Calls()) != 0")
	}
	mockedMethod.Called()
	if len(mockedMethod.Calls()) != 1 {
		t.Fatal("len(mockedMethod.Calls()) != 1")
	}
}
