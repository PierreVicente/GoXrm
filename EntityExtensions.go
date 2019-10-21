package GoXrm

import (
	"fmt"
	"reflect"
	"time"

	"github.com/PierreVicente/GoXrm/Constants"
	"github.com/google/uuid"
)

const _formattedSuffix = "@OData.Community.Display.V1.FormattedValue"
const _associatedNavigationSuffix = "@Microsoft.Dynamics.CRM.associatednavigationproperty"
const _lookupLogicalnameSuffix = "@Microsoft.Dynamics.CRM.lookuplogicalname"

func (e *Entity) GetString(attribute string) string {
	o, ok := e.GetAttributeValue(attribute)
	if ok {
		return fmt.Sprintf("%v", o)
	}
	return ""
}

func (e *Entity) SetString(attribute string, value string) {
	e.Attributes[attribute] = value
}

func (e *Entity) GetInt(attribute string) int64 {
	attr, ok := e.GetAttributeValue(attribute)
	if ok {
		i, ok := attr.(int64)
		if !ok {
			return 0
		}
		return i
	}
	return 0
}

func (e *Entity) SetInt(attribute string, value int64) {
	e.Attributes[attribute] = value
}

func (e *Entity) GetEntityReference(attribute string) EntityReference {
	//id
	eref := NewEntityReference("", "")
	if value, ok := e.Contains(attribute); ok {
		eref.Id = fmt.Sprintf("%v", value)
	} else {
		eref.Id = Constants.GuidEmpty
	}
	//logicalname
	if str, ok := e.jProps[attribute+_lookupLogicalnameSuffix]; ok {
		eref.LogicalName = fmt.Sprintf("%v", str)
	}
	//name
	if str, ok := e.jProps[attribute+_formattedSuffix]; ok {
		eref.Name = fmt.Sprintf("%v", str)
	}
	return eref
}

func (e *Entity) GetEntityReferenceId(attribute string) string {
	realE := *e
	return realE.GetEntityReference(attribute).Id
}

func (e *Entity) GetEntityReferenceName(attribute string) string {
	realE := *e
	return realE.GetEntityReference(attribute).Name
}

func (e *Entity) SetEntityReference(attribute string, reference EntityReference) {

	e.Attributes[attribute] = reference.Id
	e.FormattedValues[attribute] = reference.Name
	e.jProps[attribute+_lookupLogicalnameSuffix] = reference.LogicalName
	e.jProps[attribute+_formattedSuffix] = reference.Name

}

func (e *Entity) GetOptionSetValue(attribute string) OptionSetValue {
	opt := NewOptionSetValue(0)
	if attr, ok := e.Contains(attribute); ok {
		opt.Value = attr.(int64)
	} else {
		opt.Value = -1
	}
	//logicalname
	if str, ok := e.jProps[attribute+_formattedSuffix]; ok {
		opt.Description = fmt.Sprintf("%v", str)
	}
	return opt
}

func (e *Entity) SetOptionSetValue(attribute string, option OptionSetValue) {
	e.Attributes[attribute] = option.Value
	delete(e.FormattedValues, attribute)
	delete(e.jProps, attribute+_lookupLogicalnameSuffix)
	e.jProps[attribute+_formattedSuffix] = option.Description
}

func (e *Entity) GetOptionSetInt(attribute string) int64 {
	realE := *e
	return realE.GetOptionSetValue(attribute).Value
}

func (e *Entity) GetOptionSetName(attribute string) string {
	realE := *e
	return realE.GetOptionSetValue(attribute).Description
}

func (e *Entity) GetDec(attribute string) float64 {
	if attr, ok := e.Contains(attribute); ok {
		return attr.(float64)
	}
	return 0
}

func (e *Entity) SetDec(attribute string, value float64) {
	e.Attributes[attribute] = value
	delete(e.FormattedValues, attribute)
}

func (e *Entity) GetDateTime(attribute string) time.Time {
	if attr, ok := e.Contains(attribute); ok {
		strDt := fmt.Sprintf("%v", attr)
		t, err := time.Parse(time.RFC3339, strDt)
		if err == nil {
			return t
		}
		return time.Time{}
	}
	return time.Time{}
}

func (e *Entity) SetDateTime(attribute string, value time.Time) {
	e.Attributes[attribute] = value.Format(time.RFC3339)
	delete(e.FormattedValues, attribute)
}

func (e *Entity) GetBool(attribute string) bool {
	if attr, ok := e.Contains(attribute); ok {
		return attr.(bool)
	}
	return false
}

func (e *Entity) SetBool(attribute string, value bool) {
	e.Attributes[attribute] = value
	delete(e.FormattedValues, attribute)
	delete(e.jProps, attribute+_formattedSuffix)
}

func (e *Entity) GetMoney(attribute string) Money {
	if attr, ok := e.Contains(attribute); ok {
		return NewMoney(attr.(float64))
	}
	return NewMoney(0)
}

func (e *Entity) SetMoney(attribute string, value float64) {
	e.Attributes[attribute] = value
	delete(e.FormattedValues, attribute)
	delete(e.jProps, attribute+_formattedSuffix)
}

func (e *Entity) GetGuid(attribute string) uuid.UUID {
	if attr, ok := e.Contains(attribute); ok {
		str := fmt.Sprintf("%v", attr)
		g, err := uuid.Parse(str)
		if err == nil {
			return g
		} else {
			panic(err)
		}
	}
	return uuid.Nil
}

func (e *Entity) SetGuid(attribute string, value uuid.UUID) {
	e.Attributes[attribute] = value.String()
}

func (e *Entity) GetFormattedValue(attribute string) string {
	if intf, ok := e.FormattedValues[attribute]; ok {
		return intf
	}
	return ""
}

func (e *Entity) RemoveAttribute(attribute string) {
	delete(e.Attributes, attribute)
	delete(e.FormattedValues, attribute)
	delete(e.jProps, attribute+_formattedSuffix)
	delete(e.jProps, attribute+_lookupLogicalnameSuffix)
}

func (eptr *Entity) IsEntityReference(attribute string) (bool, string) {
	e := *eptr
	refEntity, ok := e.jProps[attribute+_lookupLogicalnameSuffix]
	if ok {
		return ok, refEntity
	}

	retAttr, ok := e.Contains(attribute)
	if ok {
		if reflect.TypeOf(retAttr) == reflect.TypeOf(EntityReference{}) {
			eref, ok2 := retAttr.(EntityReference)
			if ok2 {
				return true, eref.LogicalName
			}
		}

		//if retAttr.Type == Metadata.AttributeType_Lookup {
		//	return true, retAttr.Value.(EntityReference).LogicalName
		//}
	}

	return ok, ""

}
