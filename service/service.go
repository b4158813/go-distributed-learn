package service

import (
	"context"
	"log"
	"net/http"
)

func Start(ctx context.Context, servicename, host, port string, RegisterHandlersFunc func()) (context.Context, error) {
	RegisterHandlersFunc()
	ctx = StartService(ctx, servicename, host, port)
	return ctx, nil
}

func StartService(ctx context.Context, servicename, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		log.Printf("%v started at %s:%s ...\n", servicename, host, port)
		log.Println(srv.ListenAndServe())
		cancel()
	}()
	return ctx
}
