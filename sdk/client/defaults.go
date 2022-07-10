package client

type URL string
type Method string
type Action string

const (
	BaseUrl = "https://api.contabo.com"

	AuthUrl                 = "https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token"
	ComputeInstancesUrl URL = BaseUrl + "/v1/compute/instances"

	GET    Method = "GET"
	PUT    Method = "PUT"
	POST   Method = "POST"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
	START  Action = "start"
	REBOOT Action = "restart"
	STOP   Action = "stop"
)
