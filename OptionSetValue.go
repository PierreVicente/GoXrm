package GoXrm

type OptionSetValue struct {
	Value       int64
	Description string
}

func NewOptionSetValue(value int64) OptionSetValue {
	return OptionSetValue{Value: value}
}
