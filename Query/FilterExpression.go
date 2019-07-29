package Query

type FilterExpression struct {
	FilterOperator int32
	Conditions     []ConditionExpression
	Filters        []FilterExpression
}

func NewFilterExpression(logicalOperator int32) FilterExpression {
	return new(FilterExpression{LogicalOperator: logicalOperator})
}

func (filter *FilterExpression) AddCondition(entityName string, attributename string, conditionOperator int32, values []interface{}) {
	filter.Conditions = append(filter.Conditions, new(ConditionExpression{Values: values, EntityName: entityName, AttributeName: attributename, Operator: conditionOperator}))
}
