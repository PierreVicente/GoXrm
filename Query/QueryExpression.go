package Query

type QueryExpression struct {
	EntityName     string
	Distinct       bool
	NoLock         bool
	PageInfo       PagingInfo
	LinkEntities   []LinkEntity
	Criteria       FilterExpression
	Orders         []OrderExpression
	ColumnSet      ColumnSet
	TopCount       int32
	AggregateCount bool
	AggregateSum   bool
	AggregateAvg   bool
}

func (qe *QueryExpression) AddLink(linkToEntityName string, linkFromAttributeName string, linkToAttributeName string, joinOperator int32) LinkEntity {
	lk := new(LinkEntity{LinkFromEntityName: qe.EntityName, LinkToEntityName: linkToEntityName, LinkFromAttributeName: linkFromAttributeName, LinkToAttributeName: linkFromAttributeName, JoinOperator: joinOperator})
	qe.LinkEntities = append(qe.LinkEntities, lk)
	return lk
}

func (qe *QueryExpression) AddOrder(attributeName string, OrderType int32) {
	qe.Orders = append(qe.Orders, new(OrderExpression{OrderType: OrderType, AttributeName: attributeName}))
}

func NewQueryExpression(entityName string) *QueryExpression {
	qe := new(QueryExpression)
	qe.EntityName = entityName
	return qe
}
