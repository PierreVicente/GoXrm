package GoXrm

import (
	"strings"
)

var spnames = map[string]string{
	"appointment":        "activityid",
	"campaignactivity":   "activityid",
	"phonecall":          "activityid",
	"fax":                "activityid",
	"letter":             "activityid",
	"campaignresponse":   "activityid",
	"email":              "activityid",
	"serviceappointment": "activityid",
	"activitypointer":    "activityid",
}

func getPrimaryIdAttribute(entity string) string {
	pk, ok := spnames[entity]
	if ok {
		return pk
	} else {
		return entity + "id"
	}
}

func getCollectionSchemaName(entity string) string {
	if strings.HasSuffix(entity, "y") {
		return entity[:len(entity)-1] + "ies"
	} else if strings.HasSuffix(entity, "s") || strings.HasSuffix(entity, "x") {
		return entity + "es"
	} else {
		return entity + "s"
	}
}

func IsActivityEntity(entity string) bool {
	_, ok := spnames[entity]
	return ok
}
