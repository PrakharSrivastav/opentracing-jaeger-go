package services

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

type HTTPService struct {
}

func NewHTTPService() *HTTPService {
	return &HTTPService{}
}

func (h *HTTPService) DoSomething(ctx context.Context) {
	s, _ := opentracing.StartSpanFromContext(ctx, "http-service")
	defer s.Finish()
	s.LogFields(
		log.String("method", "DoSomethign in Http"),
		log.String("operation", fmt.Sprintf("Doing something with Http %s", time.Now().String())),
	)
	time.Sleep(time.Millisecond * 80)
}
