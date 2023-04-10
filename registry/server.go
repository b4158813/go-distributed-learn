package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type registry struct {
	Services []Registration
	mu       *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Services = append(r.Services, reg)
	return nil
}

var rsl = registry{
	Services: make([]Registration, 0),
	mu:       new(sync.Mutex),
}

type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received.")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var re Registration
		err := dec.Decode(&re)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: [%s] with URL: [%s]\n", re.ServiceName, re.ServiceUrl)
		err = rsl.add(re)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
