package GoXrm

type EntityReference struct {
	Id          string
	LogicalName string
	Name        string
}

func NewEntityreference2(LogicalName string, Id string) *EntityReference {
	eref := new(EntityReference)
	eref.LogicalName = LogicalName
	eref.Id = Id
	return eref
}
