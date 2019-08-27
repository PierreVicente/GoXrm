package Client

import "github.com/PierreVicente/GoXrm"

func (this *CrmServiceClient) Update(target GoXrm.Entity) {
	this.Upsert(target, true)
}
