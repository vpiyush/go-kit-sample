package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/vpiyush/go-kit-sample/pkg/api"
	"github.com/vpiyush/go-kit-sample/pkg/services"
	"github.com/vpiyush/go-kit-sample/pkg/storage/inmem"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	var svc services.PairService
	svc = services.NewService(inmem.NewStorage())
	svc = api.NewLoggingMiddleware(logger, svc)
	handler := api.NewHttpServer(svc)
	logger.Log("err", http.ListenAndServe(":9999", handler))
}
