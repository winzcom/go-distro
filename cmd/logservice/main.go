package main

import (
	"context"
	"fmt"
	"go-l-plurasight/log"
	"go-l-plurasight/registry"
	"go-l-plurasight/service"
	stlog "log"
)

func main() {
	log.Run("./app.log")
	host, port := "localhost", "4000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)
	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceUrl = serviceAddress

	ctx, err := service.Start(context.Background(), host, port, r,log.RegisterHandlers)
	r.RequireServices = make([]registry.ServiceName, 0)
	r.ServiceUpdateUrl = r.ServiceUrl+"/services"

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()

	fmt.Println("shutting down log service")
}
