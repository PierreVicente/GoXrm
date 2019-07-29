package GoXrm

import (
	"github.com/google/uuid"
	"go/types"
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
	Id              uuid.UUID
	LogicalName     string
	Attributes      map[string]XrmType
	FormattedValues map[string]string
	//RelatedEntities map[RelationShip]EntityCollection
	jProps []types.Var
}

func (e Entity) PrimaryIdAttribute() string {
	return GetPrimaryIdAttribute(e.LogicalName)
}

func (e Entity) GetAttributeValue(attributeName string) (XrmType, bool) {
	return e.Contains(attributeName)
}

func (e Entity) Contains(attributeName string) (XrmType, bool) {
	if val, ok := e.Attributes[attributeName]; ok {
		return val, true
	}
	if val, ok := e.Attributes[strings.ToLower(attributeName)]; ok {
		return val, true
	}

	return XrmType{TypeName: "", IntValue: 0, HasValue: false, DecimalValue: 0, GuidValue: uuid.Nil, StringValue: "", DateValue: time.Time{}, LogicalName: "", Name: ""}, false
}

func NewEntity2(logicalName string, Id uuid.UUID) *Entity {
	e := new(Entity)
	e.Id = Id
	e.LogicalName = logicalName
	return e
}
