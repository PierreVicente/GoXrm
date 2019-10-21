package GoXrm

import (
	"github.com/PierreVicente/GoXrm/Constants"
	"strings"
)

//type Attribute struct {
//	Type  int `json:"-"`
//	Value interface{}
//}

type Entity struct {
	RowVersion      int64
	Id              string
	LogicalName     string
	Attributes      map[string]interface{}
	FormattedValues map[string]string
	RelatedEntities map[RelationShip]EntityCollection
	jProps          map[string]string
}

func (e *Entity) PrimaryIdAttribute() string {
	return GetPrimaryIdAttribute(e.LogicalName)
}

func (e *Entity) GetAttributeValue(attributeName string) (interface{}, bool) {
	return e.Contains(attributeName)
}

func (e *Entity) Contains(attributeName string) (interface{}, bool) {
	if val, ok := e.Attributes[attributeName]; ok {
		return val, true
	}
	if val, ok := e.Attributes[strings.ToLower(attributeName)]; ok {
		return val, true
	}
	return nil, false
}

func NewEntity(logicalName string, id string) Entity {
	e := Entity{LogicalName: logicalName}
	e.Attributes = make(map[string]interface{})
	e.jProps = make(map[string]string)
	e.FormattedValues = make(map[string]string)
	if id == "" {
		e.Id = Constants.GuidEmpty
	}
	return e
}

func (e *Entity) ToEntityReference() EntityReference {
	return NewEntityReference(e.LogicalName, e.Id)
}
