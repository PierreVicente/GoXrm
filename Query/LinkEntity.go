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

//func NewLinkEntityCollection(links []LinkEntity) []LinkEntity{
//	var lks []LinkEntity
//	for i, lk := range(links){
//		lks[i] = lk
//	}
//	return lks
//}

func NewLinkEntity(linkFromEntityName string, linkToEntityName string, linkFromAttributeName string, linkToAttributeName string, joinOperator int32) LinkEntity {
	var lk LinkEntity
	lk.JoinOperator = joinOperator
	lk.LinkToAttributeName = linkToAttributeName
	lk.LinkFromAttributeName = linkFromAttributeName
	lk.LinkToEntityName = linkToEntityName
	lk.LinkFromEntityName = linkFromEntityName
	return lk
}
