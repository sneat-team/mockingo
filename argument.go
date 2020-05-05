package mockingo

// NewArgument create a new argument to be stored for the func calls history
func NewArgument(name string, value interface{}) Argument {
	return Argument{name: name, value: value}
}

// Argument stores name & value of an argument passed to a mocked func
type Argument struct {
	name  string
	value interface{}
}

// Name of an argument passed to a mocked func
func (v *Argument) Name() string {
	return v.name
}

// Value of an argument passed to a mocked func
func (v *Argument) Value() interface{} {
	return v.value
}
