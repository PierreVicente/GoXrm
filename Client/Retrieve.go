package Client

import (
	"github.com/PierreVicente/GoXrm"
	"github.com/PierreVicente/GoXrm/Query"
)

func (this *CrmServiceClient) Retrieve(entityName, id string, colset Query.ColumnSet) (GoXrm.Entity, bool, error) {

	qe := Query.QueryExpression{
		NoLock:     true,
		EntityName: entityName,
		ColumnSet:  colset,
		Criteria: Query.FilterExpression{
			FilterOperator: Query.LogicalOperator_And,
			Conditions: []Query.ConditionExpression{
				Query.NewConditionExpressionSingleValue(entityName, GoXrm.GetPrimaryIdAttribute(entityName), Query.ConditionOperator_Equal, id),
			},
		},
	}

	ec, retErr := this.RetrieveMultiple(qe)
	if retErr != nil {
		return GoXrm.Entity{}, false, retErr
	}
	if len(ec.Entities) > 0 {
		return ec.Entities[0], true, nil
	} else {
		return GoXrm.Entity{}, false, nil
	}
}
