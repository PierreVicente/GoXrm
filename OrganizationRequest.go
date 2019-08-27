package GoXrm

type OrganizationRequest struct {
	RequestName string
	Parameters  map[string]interface{}
}

func (this *OrganizationRequest) GetRelativeUrlString() string {
	return this.RequestName + "()"
}
