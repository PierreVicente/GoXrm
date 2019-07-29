package Query

type OrderExpression struct {
	AttributeName string
	OrderType     int32
}

func NewOrderExperssion(attributeName string, orderType int32) OrderExpression {
	return new(OrderExpression{AttributeName: attributeName, OrderType: orderType})
}
