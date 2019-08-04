package Query

type OrderExpression struct {
	AttributeName string
	OrderType     int32
}

func NewOrderExpression(attributeName string, orderType int32) OrderExpression {
	var ord OrderExpression
	ord.AttributeName = attributeName
	ord.OrderType = orderType
	return ord
	//return new(OrderExpression{AttributeName: attributeName, OrderType: orderType})
}
