package Query

type FilterExpression struct {
	FilterOperator int32
	Conditions     []ConditionExpression
	Filters        []FilterExpression
}

func NewFilterExpression(logicalOperator int32) FilterExpression {
	var fil FilterExpression
	fil.FilterOperator = logicalOperator
	return fil
	//return new(FilterExpression{LogicalOperator: logicalOperator})
}

func (filter *FilterExpression) AddCondition(entityName string, attributename string, conditionOperator int32, values []interface{}) {
	var cnd ConditionExpression
	cnd.Values = values
	cnd.EntityName = entityName
	cnd.AttributeName = attributename
	cnd.Operator = conditionOperator
	filter.Conditions = append(filter.Conditions, cnd)
}
