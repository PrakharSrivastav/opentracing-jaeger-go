package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"github.com/yurishkuro/opentracing-tutorial/go/lib/tracing"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}
	helloTo := os.Args[1]
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)

	tracer, closer := tracing.Init("hello-world")
	defer closer.Close()
	span := tracer.StartSpan("say-hello")

	println(helloStr)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	span.LogKV("event", "println")
	span.Finish()
}
