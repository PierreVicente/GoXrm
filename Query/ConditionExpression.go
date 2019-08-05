package Query

type ConditionExpression struct {
	EntityName    string
	AttributeName string
	Operator      int
	Values        []interface{}
}

func NewConditionExpressionSingleValue(entityName string, attributeName string, conditionOperator int, value string) *ConditionExpression {
	c := new(ConditionExpression)
	c.EntityName = entityName
	c.AttributeName = attributeName
	c.Operator = conditionOperator
	c.Values = append(c.Values, value)
	return c
}

func NewConditionExpressionMultipleValues(entityName string, attributeName string, conditionOperator int, values []interface{}) *ConditionExpression {
	c := new(ConditionExpression)
	c.EntityName = entityName
	c.AttributeName = attributeName
	c.Operator = conditionOperator
	c.Values = values
	return c
}
