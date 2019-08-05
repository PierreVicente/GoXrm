package GoXrm

import (
	"fmt"
	"github.com/PierreVicente/GoXrm/Constants"
	"github.com/PierreVicente/GoXrm/Metadata"
	"github.com/buger/jsonparser"
	"strconv"
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

func GetPrimaryIdAttribute(entity string) string {
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

func JObjectToEntity(o []byte, entityName string) Entity {
	e := NewEntity(entityName, "")

	primaryIdAttribute := GetPrimaryIdAttribute(entityName)

	jsonparser.ObjectEach(o, func(attB []byte, dataB []byte, vt jsonparser.ValueType, offset int) error {
		attribute := string(attB)
		data := string(dataB)
		if strings.Index(attribute, "@") > -1 {
			e.jProps[attribute] = string(dataB)
			if attribute == "@odata.etag" {
				str := data[strings.LastIndex(data, "/")+1:]
				i, _ := strconv.ParseInt(str, 10, 64)
				e.RowVersion = i
			}

			if strings.Index(attribute, _formattedSuffix) > 0 {
				attribute = strings.ReplaceAll(attribute, _formattedSuffix, "")
				e.FormattedValues[attribute] = data
			}

			return nil
		}

		switch vt.String() {
		case "non-existent":
			fmt.Println("not implemented: " + vt.String() + " " + attribute)
			break
		case "string":
			fieldType := Metadata.AttributeType_Unknown

			if attribute == primaryIdAttribute {
				e.Id = data
				fieldType = Metadata.AttributeType_UniqueIdentifier
			}

			if ok, _ := e.IsEntityReference(attribute); ok {
				fieldType = Metadata.AttributeType_Lookup
			}

			e.Attributes[attribute] = Attribute{fieldType, data}

			break
		case "number":
			v, err := strconv.ParseFloat(data, 64)
			if err == nil {
				e.Attributes[attribute] = Attribute{Metadata.AttributeType_Decimal, v}
			} else {
				fmt.Println(err)
			}
			break
		case "object":
			fmt.Println("not implemented: " + vt.String() + " " + attribute)
			break
		case "array":
			fmt.Println("not implemented: " + vt.String() + " " + attribute)
			break
		case "boolean":
			v, err := strconv.ParseBool(data)
			if err == nil {
				e.Attributes[attribute] = Attribute{Metadata.AttributeType_Bool, v}
			} else {
				fmt.Println(err)
			}
			break
		case "null":
			fmt.Println("not implemented: " + vt.String() + " " + attribute)
			break
		case "unknown":
			fmt.Println("not implemented: " + vt.String() + " " + attribute)
			break
		}

		return nil
	})
	return e
}

func EntityToJObject(target Entity, action string, isActivityEntity bool) {

	jo := make(map[string]interface{})

	if action == "" {
		action = "C"
	}

	action = strings.ToUpper(action)[0:1]

	if action == "C" {
		if target.Id != Constants.GuidEmpty {
			jo[GetPrimaryIdAttribute(target.LogicalName)] = target.Id
		}
	}

	keys := make([]string, 0)
	for key, _ := range target.Attributes {
		keys = append(keys, key)
	}

	for _, attr := range keys {
		val, _ := target.Attributes[attr]
		ok, refEntityName := target.IsEntityReference(attr)
		if ok {
			suffix := ""
			if isActivityEntity {
				suffix = "_" + target.LogicalName
			}
			jo[attr+suffix+"@odata.bind"] = "/" + getCollectionSchemaName(refEntityName) + "(" + fmt.Sprintf("%v", val)
		} else {

		}
	}

}
