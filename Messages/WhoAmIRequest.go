package Messages

import (
	"github.com/PierreVicente/GoXrm"
	"github.com/PierreVicente/GoXrm/Client"
	"github.com/buger/jsonparser"
	"strings"
)

type WhoAmIRequest GoXrm.OrganizationRequest

func NewWhoAmIRequest() WhoAmIRequest {
	return WhoAmIRequest{
		RequestName: "WhoAmI",
		Parameters:  nil,
	}
}

func (this *WhoAmIRequest) GetRelativevUrlString() string {
	if this.RequestName == "" {
		this.RequestName = "WhoAmI"
	}
	return this.RequestName + "()"
}

func (this *WhoAmIRequest) Execute(service Client.CrmServiceClient) WhoAmIResponse {
	part := this.GetRelativevUrlString()
	res, _ := service.ExecuteWebApiFunction(part)
	if res == "" {
		toto := "toto"
		toto += "o"
	}

	e := GoXrm.NewOrganizationResponse()
	e.ResponseName = "WhoAmIResponse"

	jsonparser.ObjectEach([]byte(res), func(attB []byte, dataB []byte, vt jsonparser.ValueType, offset int) error {
		attribute := string(attB)
		data := string(dataB)
		if strings.Index(attribute, "@") > -1 {
			e.Jprops[attribute] = string(dataB)
			return nil
		}
		e.Jprops[attribute] = data

		return nil
	})

	return WhoAmIResponse{
		OraganizationId: e.Jprops["OrganizationId"],
		UserId:          e.Jprops["UserId"],
		BusinessUnitId:  e.Jprops["BusinessUnitId"],
	}

}
