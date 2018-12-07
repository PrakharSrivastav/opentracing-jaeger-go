package main

import (
	"context"
	"fmt"
	"github.com/PrakharSrivastav/tracing/internal/services"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/yurishkuro/opentracing-tutorial/go/lib/tracing"
	"net/http"
	"strings"
)

var dbService *services.DBService
var httpService *services.HTTPService

func init() {
	fmt.Println("Registering dbService")
	dbService = services.NewDBService()

	fmt.Println("Registering httpService")
	httpService = services.NewHTTPService()
}

func main() {
	fmt.Println("Starting")

	fmt.Println("Registering All Routes")
	http.HandleFunc("/", sayHello)

	fmt.Println("Starting http listener")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	tracer, closer := tracing.Init("home-page")
	defer closer.Close()

	span := tracer.StartSpan("visit-home-page")
	defer span.Finish()
	opentracing.SetGlobalTracer(tracer)

	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	dbService.DoSomething(ctx)

	httpService.DoSomething(ctx)
	message = "Hello " + message
	span.LogFields(
		log.String("url", "home-page"),
		log.String("message", message),
	)

	w.Write([]byte(message))
}
