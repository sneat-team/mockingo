package mockingo

import "testing"

func TestNewArgument(t *testing.T) {
	const name = "arg1"
	const val = "val1"
	argument := NewArgument(name, val)
	if argument.name != "arg1" {
		t.Fatalf("argument.name == '%v'", name)
	}
	if argument.value.(string) != val {
		t.Fatalf("argument.name == '%v'", val)
	}
}

func TestArgument_Name(t *testing.T) {
	const name = "arg1"
	argument := Argument{name: name}
	if actual := argument.Name(); actual != name {
		t.Fatalf("argument.Name() != '%v', actual: %v", name, actual)
	}
}
