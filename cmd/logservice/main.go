package main

import (
	"context"
	"fmt"
	stlog "log"
	"projects/log"
	"projects/registry"
	"projects/service"
)

func main() {
	log.Run("./distributed.log")
	ctx, err := service.Start(
		context.Background(),
		log.ServerHost,
		log.ServerPort,
		registry.Registration{
			ServiceName: "log service",
			ServiceUrl:  fmt.Sprintf("http://%s%s", log.ServerHost, log.ServerPort),
		},
		log.RegisterHttpHandler,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
}
