package gradingservice

import (
	"context"
	"fmt"
	"go-l-plurasight/grades"
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

	ctx, err := service.Start(context.Background(), host, port, r, grades.RegisterHandlers)

	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("shtting down grading service")
}
