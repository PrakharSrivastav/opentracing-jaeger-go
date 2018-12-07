package services

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

type DBService struct {
}

func NewDBService() *DBService {
	return &DBService{}
}

func (d *DBService) DoSomething(ctx context.Context) {
	s, _ := opentracing.StartSpanFromContext(ctx, "db-service")
	defer s.Finish()
	s.LogFields(
		log.String("method", "DoSomethign in db"),
		log.String("operation", fmt.Sprintf("Doing something with DB %s", time.Now().String())),
	)
	time.Sleep(time.Millisecond * 500)
}
