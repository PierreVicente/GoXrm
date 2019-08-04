package GoXrm

type Money struct {
	Value float64
}

func NewMoney(value float64) Money {
	return Money{Value: value}
}
