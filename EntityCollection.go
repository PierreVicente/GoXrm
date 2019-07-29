package GoXrm

type EntityCollection struct {
	EntityName                    string
	Entities                      []Entity
	MoreRecords                   bool
	PagingCookie                  string
	NextPage                      int32
	TotalRecordCount              int32
	TotalRecordCountLimitExceeded bool
}

func NewEntityCollection0() *EntityCollection {
	e := new(EntityCollection)
	e.Entities = make([]Entity, 0)
	return e
}

func NewEntityCollection2(entityName string) *EntityCollection {
	e := NewEntityCollection0()
	e.EntityName = entityName
	return e
}
