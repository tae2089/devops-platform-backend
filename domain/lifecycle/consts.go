package lifecycle

type Type string

const (
	Production  Type = "PRODUCTION"
	Development Type = "DEVELOPMENT"
	Other       Type = "OTHER"
	Deprecated  Type = "DEPRECATED"
)
