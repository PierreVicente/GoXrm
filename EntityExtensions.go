package GoXrm

import (
	"fmt"
	"github.com/PierreVicente/GoXrm/Constants"
	"github.com/google/uuid"
	"time"
)

const _formattedSuffix = "@OData.Community.Display.V1.FormattedValue"
const _associatedNavigationSuffix = "@Microsoft.Dynamics.CRM.associatednavigationproperty"
const _lookupLogicalnameSuffix = "@Microsoft.Dynamics.CRM.lookuplogicalname"

func GetString(e *Entity, attribute string) string {
	o, ok := e.GetAttributeValue(attribute)
	if ok {
		return fmt.Sprintf("%v", o)
	}
	return ""
}

func SetString(e *Entity, attribute string, value string) {
	e.Attributes[attribute] = value
}

func GetInt(e *Entity, attribute string) int64 {
	o, ok := e.GetAttributeValue(attribute)
	if ok {
		i, ok := o.(int64)
		if !ok {
			return 0
		}
		return i
	}
	return 0
}

func SetInt(e *Entity, attribute string, value int64) {
	e.Attributes[attribute] = value
}

func GetEntityReference(e *Entity, attribute string) EntityReference {
	//id
	eref := NewEntityReference("", "")
	if str, ok := e.Contains(attribute); ok {
		eref.Id = fmt.Sprintf("%v", str)
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

func GetEntityReferenceId(e *Entity, attribute string) string {
	if intf, ok := e.Contains(attribute); ok {
		return fmt.Sprintf("%v", intf)
	}
	return Constants.GuidEmpty
}

func GetEntityReferenceName(e *Entity, attribute string) string {
	if intf, ok := e.jProps[attribute+_formattedSuffix]; ok {
		return fmt.Sprintf("%v", intf)
	}
	return ""
}

func SetEntityReference(e *Entity, attribute string, reference EntityReference) {

	e.Attributes[attribute] = reference.Id
	e.FormattedValues[attribute] = reference.Name
	e.jProps[attribute] = reference.Id
	e.jProps[attribute+_lookupLogicalnameSuffix] = reference.LogicalName
	e.jProps[attribute+_formattedSuffix] = reference.Name

}

func GetOptionSetValue(e *Entity, attribute string) OptionSetValue {
	opt := NewOptionSetValue(0)
	if str, ok := e.Contains(attribute); ok {
		opt.Value = str.(int64)
	}
	//logicalname
	if str, ok := e.jProps[attribute+_formattedSuffix]; ok {
		opt.Description = fmt.Sprintf("%v", str)
	}
	return opt
}

func SetOptionSetValue(e *Entity, attribute string, option OptionSetValue) {
	e.Attributes[attribute] = option.Value
	delete(e.FormattedValues, attribute)
	e.jProps[attribute] = option.Value
	delete(e.jProps, attribute+_lookupLogicalnameSuffix)
	e.jProps[attribute+_formattedSuffix] = option.Description
}

func GetOptionSetInt(e *Entity, attribute string) int64 {
	if str, ok := e.jProps[attribute+_formattedSuffix]; ok {
		return str.(int64)
	}
	return -1
}

func GetOptionSetName(e *Entity, attribute string) string {
	if str, ok := e.jProps[attribute+_formattedSuffix]; ok {
		return fmt.Sprintf("%v", str)
	}
	return ""
}

func GetDec(e *Entity, attribute string) float64 {
	if intf, ok := e.Contains(attribute); ok {
		return intf.(float64)
	}
	return 0
}

func SetDec(e *Entity, attribute string, value float64) {
	e.Attributes[attribute] = value
	delete(e.FormattedValues, attribute)
}

func GetDateTime(e *Entity, attribute string) time.Time {
	if intf, ok := e.Contains(attribute); ok {
		strDt := fmt.Sprintf("%v", intf)
		t, err := time.Parse(time.RFC3339, strDt)
		if err == nil {
			return t
		}
		return time.Time{}
	}
	return time.Time{}
}

func SetDateTime(e *Entity, attribute string, value time.Time) {
	e.Attributes[attribute] = value.Format(time.RFC3339)
	delete(e.FormattedValues, attribute)
}

func GetBool(e *Entity, attribute string) bool {
	if intf, ok := e.Contains(attribute); ok {
		return intf.(bool)
	}
	return false
}

func SetBool(e *Entity, attribute string, value bool) {
	e.Attributes[attribute] = value
	delete(e.FormattedValues, attribute)
	delete(e.jProps, attribute+_formattedSuffix)
}

func GetMoney(e *Entity, attribute string) Money {
	if intf, ok := e.Contains(attribute); ok {
		return NewMoney(intf.(float64))
	}
	return NewMoney(0)
}

func SetMoney(e *Entity, attribute string, value float64) {
	e.Attributes[attribute] = value
	delete(e.FormattedValues, attribute)
	delete(e.jProps, attribute+_formattedSuffix)
}

func GetGuid(e *Entity, attribute string) uuid.UUID {
	if intf, ok := e.Contains(attribute); ok {
		str := fmt.Sprintf("%v", intf)
		g, err := uuid.Parse(str)
		if err == nil {
			return g
		} else {
			panic(err)
		}
	}
	return uuid.Nil
}

func SetGuid(e *Entity, attribute string, value uuid.UUID) {
	e.Attributes[attribute] = value.String()
}

func GetFormattedValue(e *Entity, attribute string) string {
	if intf, ok := e.FormattedValues[attribute]; ok {
		return intf
	}
	return ""
}

func RemoveAttribute(e *Entity, attribute string) {
	delete(e.Attributes, attribute)
	delete(e.FormattedValues, attribute)
	delete(e.jProps, attribute+_formattedSuffix)
	delete(e.jProps, attribute+_lookupLogicalnameSuffix)
}
