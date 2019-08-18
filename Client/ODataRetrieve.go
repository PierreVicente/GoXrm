package Client

import (
	"bytes"
	"github.com/PierreVicente/GoXrm"
	"github.com/PierreVicente/GoXrm/AADAuth"
	"github.com/PierreVicente/GoXrm/Query"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

func (this *CrmServiceClient) ODataRetrieve(entityName, id string, colset Query.ColumnSet) (GoXrm.Entity, bool) {
	baseUrl := this.resourceUrl
	relativeUrl := "/api/data/v" + this.ApiVersion + "/" + getCollectionSchemaName(entityName) + "(" + id + ")"

	if !colset.AllColumns && len(colset.Colmuns) > 0 {
		relativeUrl += "?$select=" + strings.Join(colset.Colmuns, ",")
	}

	req, err := http.NewRequest("GET", baseUrl+relativeUrl, bytes.NewBuffer([]byte("")))

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
	r1 := string(rbody)
	r1 = strings.ReplaceAll(r1, "\\\"", "") //Pour les valeurs odata.etag

	//var jo interface{}
	//err2 := json.Unmarshal([]byte(r1), jo)
	return GoXrm.JObjectToEntity([]byte(r1), entityName), true
}
