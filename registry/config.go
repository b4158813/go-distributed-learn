package registry

const (
	ServerHost   = "127.0.0.1"
	ServerPort   = ":3000"
	ServerRouter = "/services"
	ServiceUrl   = "http://" + ServerHost + ServerPort + ServerRouter
)
