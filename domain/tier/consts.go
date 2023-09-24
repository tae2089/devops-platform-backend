package tier

type Type string

const (
	InternalService Type = "INTERNAL_SERVICE"
	CustomerFacing  Type = "CUSTOMER_FACING"
	Other           Type = "OTHER"
	MissionCritical Type = "MISSION_CRITICAL"
)
