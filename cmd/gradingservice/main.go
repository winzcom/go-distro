package gradingservice

import (
	"context"
	"fmt"
	"go-l-plurasight/grades"
	"go-l-plurasight/log"
	"go-l-plurasight/registry"
	"go-l-plurasight/service"
	stlog "log"
)

func main() {
	host, port := "localhost", "6000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration

	r.ServiceName = registry.GradingService
	r.ServiceUrl = serviceAddress
	r.RequireServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateUrl = r.ServiceUrl+"/services"

	ctx, err := service.Start(context.Background(), host, port, r, grades.RegisterHandlers)

	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProviders(registry.LogService); err ==  nil {
		fmt.Printf("logging service was found at %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("shtting down grading service")
}
