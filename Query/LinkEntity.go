package Query

type LinkEntity struct {
	LinkFromAttributeName string
	LinkFromEntityName    string
	LinkToAttributeName   string
	LinkToEntityName      string
	LinkCriteria          FilterExpression
	LinkEntities          []LinkEntity
	Columns               ColumnSet
	EntityAlias           string
	Orders                []OrderExpression
	JoinOperator          int32
}
