package GoXrm

import (
	"github.com/buger/jsonparser"
)

type EntityCollection struct {
	EntityName                    string
	Entities                      []Entity
	MoreRecords                   bool
	PagingCookie                  string
	NextPage                      int64
	TotalRecordCount              int64
	TotalRecordCountLimitExceeded bool
}

func NewEntityCollection0() EntityCollection {
	e := EntityCollection{}
	e.Entities = make([]Entity, 0)
	return e
}

func NewEntityCollection2(entityName string) EntityCollection {
	e := NewEntityCollection0()
	e.EntityName = entityName
	return e
}

func (ec *EntityCollection) FillEntities(arr []byte) {

	jsonparser.ArrayEach(arr, func(o []byte, dataType jsonparser.ValueType, offset int, err error) {
		ec.Entities = append(ec.Entities, JObjectToEntity(o, ec.EntityName))
	})
}
