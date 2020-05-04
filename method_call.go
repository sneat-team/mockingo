package mockingo

type MethodCall struct {
	args map[string]interface{}
}

func (v *MethodCall) Args() map[string]interface{} {
	return v.args
}
