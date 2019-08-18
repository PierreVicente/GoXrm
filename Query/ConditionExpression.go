package Query

type ConditionExpression struct {
	EntityName    string
	AttributeName string
	Operator      int
	Values        []interface{}
}

func NewConditionExpressionSingleValue(entityName string, attributeName string, conditionOperator int, value string) ConditionExpression {
	c := ConditionExpression{
		EntityName:    entityName,
		AttributeName: attributeName,
		Operator:      conditionOperator,
	}
	c.Values = append(make([]interface{}, 0), value)
	return c
}

func NewConditionExpressionMultipleValues(entityName string, attributeName string, conditionOperator int, values []interface{}) ConditionExpression {
	c := ConditionExpression{
		EntityName:    entityName,
		AttributeName: attributeName,
		Operator:      conditionOperator,
		Values:        values,
	}
	return c
}
