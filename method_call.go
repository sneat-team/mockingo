package mockingo

// MethodCall stores information about a call to a mocked func
type MethodCall struct {
	args map[string]interface{}
}

// Args passed with a call to a mocked func
func (v *MethodCall) Args() map[string]interface{} {
	return v.args
}
