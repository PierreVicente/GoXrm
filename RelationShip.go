package GoXrm

const (
	EntityRole_Referenced  = 1
	EntityRole_Referencing = 2
)

type RelationShip struct {
	EntityRole int32
	SchemaName string
}
