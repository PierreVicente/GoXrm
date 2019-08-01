package Client

import (
	"/"
	"github.com/google/uuid"
)

type CrmServiceClient struct {
	CallerId   uuid.UUID
	ApiVersion string
	_auth      AADAuth.AADAuthResult
}
