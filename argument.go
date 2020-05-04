package mockingo

func NewArgument(name string, value interface{}) Argument {
	return Argument{name: name, value: value}
}

type Argument struct {
	name  string
	value interface{}
}

func (v *Argument) Name() string {
	return v.name
}

func (v *Argument) Value() interface{} {
	return v.value
}
