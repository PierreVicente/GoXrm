package GoXrm

import (
	"github.com/PierreVicente/GoXrm/Constants"
	"github.com/google/uuid"
	"strings"
	"time"
)

type XrmType struct {
	TypeName     string
	LogicalName  string
	Name         string
	IntValue     int64
	StringValue  string
	DateValue    time.Time
	GuidValue    uuid.UUID
	DecimalValue float64
	HasValue     bool
}

func (xrmt XrmType) IsNull() bool {
	if xrmt.DateValue.IsZero() && xrmt.StringValue == "" && xrmt.GuidValue == uuid.Nil && xrmt.IntValue == 0 && xrmt.DecimalValue == 0 && !xrmt.HasValue {
		return true
	}
	return false
}

type Entity struct {
	RowVersion      int64
	Id              string
	LogicalName     string
	Attributes      map[string]interface{}
	FormattedValues map[string]string
	//RelatedEntities map[RelationShip]EntityCollection
	jProps map[string]interface{}
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

	return XrmType{TypeName: "", IntValue: 0, HasValue: false, DecimalValue: 0, GuidValue: uuid.Nil, StringValue: "", DateValue: time.Time{}, LogicalName: "", Name: ""}, false
}

func NewEntity(logicalName string, id string) Entity {
	e := Entity{LogicalName: logicalName}
	e.Attributes = make(map[string]interface{})
	e.jProps = make(map[string]interface{})
	e.FormattedValues = make(map[string]string)
	if id == "" {
		e.Id = Constants.GuidEmpty
	}
	return e
}
