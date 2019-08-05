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
	"strconv"
	"strings"
)

func (this *CrmServiceClient) ODataRetrieveMultiple(query Query.QueryExpression) GoXrm.EntityCollection {

	res := GoXrm.EntityCollection{EntityName: query.EntityName}

	if query.PageInfo.Count == 0 || query.PageInfo.PageNumber == 0 {
		query.PageInfo = Query.PagingInfo{Count: 5000, PageNumber: 1}
	}

	baseUrl := this.aadAuthResult.Resource

	fullUrl := ""
	if query.PageInfo.PagingCookie == "" {
		qs := query.ToODataQuery("")
		relativeUrl := "/api/data/v" + this.ApiVersion + "/" + getCollectionSchemaName(query.EntityName) + "?" + qs
		fullUrl = baseUrl + relativeUrl
	} else {
		fullUrl = query.PageInfo.PagingCookie
	}

	var stringContent = []byte("")

	req, err := http.NewRequest("GET", fullUrl, bytes.NewBuffer(stringContent))

	if err != nil {
		panic(err)
	}
	if this.aadAuthResult.AuthType != AADAuth.AuthType_AD {
		req.Header.Add("Authorization", "Bearer "+this.aadAuthResult.AccessToken)
	}
	req.Header.Add("OData-MaxVersion", "4.0")
	req.Header.Add("OData-Version", "4.0")
	req.Header.Add("Accept", "application/json")
	prefer := "odata.include-annotations=*"
	if query.PageInfo.Count != 5000 && query.PageInfo.Count != 0 {
		prefer += ";" + "odata.maxpagesize=" + fmt.Sprintf("%v", query.PageInfo.Count)
	}
	req.Header.Add("Prefer", prefer)

	if this.CallerId != uuid.Nil {
		req.Header.Add("MSCRMCallerID", this.CallerId.String())
	}

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

	jsonNodebArr, _, _, err4 := jsonparser.Get(bjs, "@odata.nextLink")
	if err4 != nil {
		panic(err4)
	}
	pagingcookie := string(jsonNodebArr)

	if pagingcookie != "" {
		res.MoreRecords = true
		res.PagingCookie = pagingcookie
	} else {
		res.MoreRecords = false
		res.PagingCookie = ""
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
