package registry

const (
	ServerHost   = "localhost"
	ServerPort   = ":3000"
	ServerRouter = "/services"
	ServiceUrl   = "http://" + ServerHost + ServerPort + ServerRouter
)
