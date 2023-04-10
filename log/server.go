package log

import (
	"io"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func Run(dest string) {
	log = stlog.New(fileLog(dest), "MyProj: ", stlog.LstdFlags)
}

func RegisterHttpHandler() {
	http.HandleFunc(ServerRouter, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Printf("%v\n", string(msg))
		default:
			w.Write([]byte("not allowed."))
			// w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}
