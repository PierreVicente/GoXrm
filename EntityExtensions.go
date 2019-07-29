package GoXrm

const _formattedSuffix = "@OData.Community.Display.V1.FormattedValue"
const _associatedNavigationSuffix = "@Microsoft.Dynamics.CRM.associatednavigationproperty"
const _lookukLogicalnameSuffix = "@Microsoft.Dynamics.CRM.lookuplogicalname"

func GetString(e *Entity, attribute string) string {
	o, ok := e.GetAttributeValue(attribute)
	if ok {
		return o.StringValue
	}
	return ""
}

func GetInt(e *Entity, attribute string) int64 {
	o, ok := e.GetAttributeValue(attribute)
	if ok {
		return o.IntValue
	}
	return 0
}
