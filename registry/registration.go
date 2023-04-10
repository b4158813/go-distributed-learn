package registry

type Registration struct {
	ServiceName string `json:"servicename"`
	ServiceUrl  string `json:"serviceurl"`
}

const (
	LogService = "LogService"
)
