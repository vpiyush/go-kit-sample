package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/vpiyush/go-kit-sample/pkg/services"
)

type (
	addPairRequest struct {
		Key   string `validate:"required" json:"key"`
		Value string `validate:"required" json:"value"`
	}
	addPairResponse struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	getPairRequest struct {
		Key string `validate:"required" json:"key"`
	}
	getPairResponse struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
)

func makeAddPairEndpoint(s services.PairService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addPairRequest)
		v, err := s.Insert(req.Key, req.Value)
		if err != nil {
			return addPairResponse{"", ""}, err
		}
		return addPairResponse{v.Key, v.Value}, nil
	}
}
func makeGetPairEndpoint(s services.PairService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getPairRequest)
		v, err := s.Get(req.Key)
		if err != nil {
			return getPairResponse{"", ""}, err
		}
		return getPairResponse{v.Key, v.Value}, nil
	}
}
