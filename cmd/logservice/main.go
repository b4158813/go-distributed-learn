package main

import (
	"context"
	stlog "log"
	"projects/log"
	"projects/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "123.60.70.187", "4000"
	ctx, err := service.Start(
		context.Background(),
		"Log service",
		host,
		port,
		log.RegisterHttpHandler,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
}
