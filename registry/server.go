package registry

import "sync"

type RegisteredServiceSt struct {
	Services []ServiceRegistrationSt
	mu       *sync.Mutex
}

var rsl = RegisteredServiceSt{
	Services: make([]ServiceRegistrationSt, 0),
	mu:       new(sync.Mutex),
}
