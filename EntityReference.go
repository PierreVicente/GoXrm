package GoXrm

import "github.com/PierreVicente/GoXrm/Constants"

type EntityReference struct {
	Id          string
	LogicalName string
	Name        string
}

func NewEntityReference(logicalName string, id string) EntityReference {
	eref := EntityReference{Id: id, LogicalName: logicalName}
	return eref
}

func IsNull(eref *EntityReference) bool {
	return (eref.Id == "" || eref.Id == Constants.GuidEmpty) && eref.LogicalName == "" && eref.Name == ""
}
