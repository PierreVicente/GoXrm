package Client

import (
	"bytes"
	"github.com/PierreVicente/GoXrm"
	"github.com/PierreVicente/GoXrm/AADAuth"
	"github.com/google/uuid"
	"net/http"
)

func (this *CrmServiceClient) Upsert(target GoXrm.Entity, updateOnly bool) {
	relativeUrl := "/api/data/v" + this.ApiVersion + "/" + getCollectionSchemaName(target.LogicalName)
	baseUrl := this.aadAuthResult.Resource

	jo := GoXrm.EntityToJObject(target, "U", false)

	body := jo
	stringContent := []byte(body)
	req, err := http.NewRequest("PATCH", baseUrl+relativeUrl, bytes.NewBuffer(stringContent))

	if err != nil {
		panic(err)
	}
	if this.aadAuthResult.AuthType != AADAuth.AuthType_AD {
		req.Header.Add("Authorization", "Bearer "+this.aadAuthResult.AccessToken)
	}
	req.Header.Add("OData-MaxVersion", "4.0")
	req.Header.Add("OData-Version", "4.0")

	if this.CallerId != uuid.Nil {
		req.Header.Add("MSCRMCallerID", this.CallerId.String())
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	if updateOnly {
		req.Header.Add("If-Match", "*")
	}

	client := &http.Client{}
	var resp, err2 = client.Do(req)
	if err2 != nil {
		panic(err2)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {

	} else {
		panic("marchepas")
	}
}
