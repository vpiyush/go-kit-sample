package api

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/vpiyush/go-kit-sample/pkg/services"
	"github.com/vpiyush/go-kit-sample/pkg/storage/inmem"
	"net/http"
)

func NewHttpServer() *mux.Router {
	svc := services.NewService(inmem.NewStorage())
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	addPairHandler := httptransport.NewServer(
		makeAddPairEndpoint(svc),
		decodeAddPairRequest,
		encodeResponse)
	getPairHandler := httptransport.NewServer(
		makeGetPairEndpoint(svc),
		decodeGetPairRequest,
		encodeResponse)

	r.Methods("POST").Path("/pair").Handler(addPairHandler)
	r.Methods("GET").Path("/pair/{key}").Handler(getPairHandler)
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("content-type", "application/json")
		next.ServeHTTP(w, req)
	})
}
