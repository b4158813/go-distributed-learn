package log

const (
	ServerHost   = "127.0.0.1"
	ServerPort   = ":4000"
	ServerRouter = "/log"
	ServiceUrl   = "http://" + ServerHost + ServerPort + ServerRouter
)
