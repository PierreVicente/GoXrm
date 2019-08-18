package Client

import (
	"bytes"
	"github.com/PierreVicente/GoXrm/AADAuth"
	"github.com/google/uuid"
	"net/http"
)

func (this *CrmServiceClient) Delete(entityLogicalName, id string) bool {
	relativeUrl := "/api/data/v" + this.ApiVersion + "/" + getCollectionSchemaName(entityLogicalName)
	relativeUrl += "(" + id + ")"
	baseUrl := this.resourceUrl

	req, err := http.NewRequest("DELETE", baseUrl+relativeUrl, bytes.NewBuffer([]byte("")))

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
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	var resp, err2 = client.Do(req)
	if err2 != nil {
		panic(err2)
	}
	defer resp.Body.Close()

	return (resp.StatusCode >= 200 && resp.StatusCode < 300)
}
