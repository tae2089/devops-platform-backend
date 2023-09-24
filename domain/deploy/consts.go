package deploy

type Kind string

const (
	Deployment  Kind = "Deployment"
	CronJob     Kind = "CronJob"
	Job         Kind = "Job"
	StatefulSet Kind = "StatefulSet"
	DaemonSet   Kind = "DaemonSet"
	OnDemand    Kind = "OnDemand"
)
