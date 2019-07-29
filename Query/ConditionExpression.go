package Query

type ConditionExpression struct {
	EntityName    string
	AttributeName string
	Operator      int32
	Values        []interface{}
}

func NewConditionExpressionSingleValue(entityName string, attributeName string, conditionOperator int32, value interface{}) *ConditionExpression {
	c := new(ConditionExpression)
	c.EntityName = entityName
	c.AttributeName = attributeName
	c.Operator = conditionOperator
	c.Values = append(c.Values, value)
	return c
}

func NewConditionExpressionMultipleValues(entityName string, attributeName string, conditionOperator int32, values []interface{}) *ConditionExpression {
	c := new(ConditionExpression)
	c.EntityName = entityName
	c.AttributeName = attributeName
	c.Operator = conditionOperator
	c.Values = values
	return c
}
