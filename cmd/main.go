package main

import (
	"github.com/vpiyush/go-kit-sample/pkg/api"
	"log"
	"net/http"
)

func main() {
	handler := api.NewHttpServer()
	log.Fatal(http.ListenAndServe(":9999", handler))
}
