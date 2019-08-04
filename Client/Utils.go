package Client

import "strings"

func getCollectionSchemaName(entity string) string {

	if strings.HasSuffix(entity, "y") {
		return entity[0:len(entity)-1] + "ies"
	} else if strings.HasSuffix(entity, "s") || strings.HasSuffix(entity, "x") {
		return entity + "es"
	} else {
		return entity + "s"
	}

}
