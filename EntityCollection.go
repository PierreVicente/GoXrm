package GoXrm

import (
	"fmt"
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
)

type EntityCollection struct {
	EntityName                    string
	Entities                      []Entity
	MoreRecords                   bool
	PagingCookie                  string
	NextPage                      int64
	TotalRecordCount              int64
	TotalRecordCountLimitExceeded bool
}

func NewEntityCollection0() *EntityCollection {
	e := new(EntityCollection)
	e.Entities = make([]Entity, 0)
	return e
}

func NewEntityCollection2(entityName string) *EntityCollection {
	e := NewEntityCollection0()
	e.EntityName = entityName
	return e
}

func (ec *EntityCollection) FillEntities(arr []byte) {

	jsonparser.ArrayEach(arr, func(o []byte, dataType jsonparser.ValueType, offset int, err error) {
		ec.Entities = append(ec.Entities, JObjectToEntity(o, ec.EntityName))
	})
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
			e.Attributes[attribute] = data

			if attribute == primaryIdAttribute {
				e.Id = data
			}

			break
		case "number":
			v, err := strconv.ParseFloat(data, 64)
			if err == nil {
				e.Attributes[attribute] = v
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
				e.Attributes[attribute] = v
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
