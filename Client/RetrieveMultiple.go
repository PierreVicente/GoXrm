package Client

import (
	"bytes"
	"fmt"
	"github.com/PierreVicente/GoXrm"
	"github.com/PierreVicente/GoXrm/AADAuth"
	"github.com/PierreVicente/GoXrm/Query"
	"github.com/buger/jsonparser"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type RMRaw struct {
	ODataContext                  string      `json "@odata.context"`
	TotalRecordCount              int64       `json "@Microsoft.Dynamics.CRM.totalrecordcount"`
	TotalRecordCountLimitExceeded bool        `json "@Microsoft.Dynamics.CRM.totalrecordcountlimitexceeded"`
	PagingCookie                  string      `json "@Microsoft.Dynamics.CRM.fetchxmlpagingcookie"`
	MoreRecords                   bool        `json "@Microsoft.Dynamics.CRM.morerecords"`
	Value                         interface{} `json "value"`
}

func (this *CrmServiceClient) RetrieveMultiple(query Query.QueryExpression) GoXrm.EntityCollection {

	res := GoXrm.EntityCollection{EntityName: query.EntityName}

	if query.PageInfo.Count == 0 || query.PageInfo.PageNumber == 0 {
		query.PageInfo = Query.PagingInfo{Count: 5000, PageNumber: 1}
	}

	q := query.ToFetchXml()
	eFetchxml := url.QueryEscape(q)
	baseUrl := this.GetAadAuthResult().Resource
	relativeUrl := "/api/data/v" + this.ApiVersion

	body := "--batch_recordfetch\n"
	body += "Content-Type: application/http\n"
	body += "Content-Transfer-Encoding: binary\n"
	body += "\n"
	body += "GET " + baseUrl + relativeUrl + "/" + getCollectionSchemaName(query.EntityName) + "?fetchXml=" +
		eFetchxml + " HTTP/1.1\n"
	body += "Content-Type: application/json\n"
	body += "OData-Version: 4.0\n"
	body += "OData-MaxVersion: 4.0\n"
	body += "Prefer: odata.include-annotations=*\n"
	body += "\n"
	body += "--batch_recordfetch--"

	var stringContent = []byte(body)

	req, err := http.NewRequest("POST", baseUrl+relativeUrl+"/$batch", bytes.NewBuffer(stringContent))

	if err != nil {
		panic(err)
	}
	if this.aadAuthResult.AuthType != AADAuth.AuthType_AD {
		req.Header.Add("Authorization", "Bearer "+this.aadAuthResult.AccessToken)
	}
	req.Header.Add("OData-MaxVersion", "4.0")
	req.Header.Add("OData-Version", "4.0")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Prefer", "odata.include-annotations=*")
	req.Header.Add("Media-Type", "multipart/mixed")
	if this.CallerId != uuid.Nil {
		req.Header.Add("MSCRMCallerID", this.CallerId.String())
	}
	req.Header.Add("Content-Type", "multipart/mixed;boundary=batch_recordfetch")

	client := &http.Client{}
	var resp, err2 = client.Do(req)
	if err2 != nil {
		panic(err2)
	}
	defer resp.Body.Close()

	rbody, _ := ioutil.ReadAll(resp.Body)
	//v:= RMRaw{}

	r1 := string(rbody)
	r1 = strings.ReplaceAll(r1, "\\\"", "")
	//r1, e := url.PathUnescape(r1)
	//if (e != nil){
	//}
	//r1 = strings.ReplaceAll(r1, "_x002e_", ".")

	p1 := strings.Index(r1, "{")
	p2 := strings.LastIndex(r1, "}") + 1

	r1 = r1[p1:p2]

	bjs := []byte(r1)

	jsonNodebArr, _, _, err4 := jsonparser.Get(bjs, "@Microsoft.Dynamics.CRM.fetchxmlpagingcookie")
	if err4 != nil {
		panic(err4)
	}
	pagingcookie := string(jsonNodebArr)

	if pagingcookie != "" {
		parts := strings.Split(pagingcookie, " ")
		for _, stmp := range parts {
			if strings.HasPrefix(stmp, "pagenumber") {
				w := strings.Split(stmp, "=")
				res.NextPage, _ = strconv.ParseInt(w[1], 10, 64)
			} else if strings.HasPrefix(stmp, "pagingcookie") {
				w := strings.Split(stmp, "=")
				pc := w[1]
				pc = strings.ReplaceAll(pc, "%25", "%")
				pc, _ = url.QueryUnescape(pc)
				res.PagingCookie = pc
			}

		}
		res.MoreRecords = true
	}
	jsonNodebArr = nil

	jsonNodebArr, _, _, err4 = jsonparser.Get(bjs, "@Microsoft.Dynamics.CRM.totalrecordcount")
	if err4 != nil {
		panic(err4)
	}
	if len(jsonNodebArr) > 0 {
		res.TotalRecordCount, _ = strconv.ParseInt(string(jsonNodebArr), 10, 64)
	}

	jsonNodebArr = nil

	jsonNodebArr, _, _, err4 = jsonparser.Get(bjs, "@Microsoft.Dynamics.CRM.totalrecordcountlimitexceeded")
	if err4 != nil {
		panic(err4)
	}
	if len(jsonNodebArr) > 0 {
		b1, err6 := strconv.ParseBool(string(jsonNodebArr))
		if err6 != nil {
			fmt.Println(err6)
		}
		res.TotalRecordCountLimitExceeded = b1 //, _ = strconv.ParseBool(string(jsonNodebArr))
	}

	jsonNodebArr, _, _, err4 = jsonparser.Get(bjs, "value")

	res.FillEntities(jsonNodebArr)

	return res
}

func (this *CrmServiceClient) RetrieveAllRecords(query Query.QueryExpression) GoXrm.EntityCollection {
	res := GoXrm.NewEntityCollection2(query.EntityName)
	var ec GoXrm.EntityCollection

	for {
		ec = this.RetrieveMultiple(query)
		for _, e := range ec.Entities {
			res.Entities = append(res.Entities, e)
		}
		if !ec.MoreRecords {
			break
		}
	}
	return res
}

func (this *CrmServiceClient) RetrieveWithRM(logicalName string, id string, columnSet Query.ColumnSet) GoXrm.Entity {

	qe := Query.QueryExpression{
		EntityName: logicalName,
		NoLock:     true,
		ColumnSet:  columnSet,
		Criteria: Query.FilterExpression{
			FilterOperator: Query.LogicalOperator_And,
			Conditions: []Query.ConditionExpression{
				Query.ConditionExpression{
					AttributeName: GoXrm.GetPrimaryIdAttribute(logicalName),
					Operator:      Query.ConditionOperator_Equal,
					Values:        []interface{}{id},
				},
			},
		},
	}

	ec := this.RetrieveMultiple(qe)
	if len(ec.Entities) > 0 {
		return ec.Entities[0]
	}

	return GoXrm.Entity{}
}
