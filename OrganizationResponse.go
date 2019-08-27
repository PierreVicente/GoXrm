package GoXrm

type OrganizationResponse struct {
	ResponseName  string
	RequestResult string
	Jprops        map[string]string
}

func NewOrganizationResponse() OrganizationResponse {
	return OrganizationResponse{
		ResponseName:  "",
		RequestResult: "",
		Jprops:        make(map[string]string),
	}
}
