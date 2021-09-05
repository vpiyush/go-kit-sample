package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeAddPairRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addPairRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	fmt.Println("request received", request)
	return request, nil
}

func decodeGetPairRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getPairRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// encode the response back,
func encodeResponse(_ context.Context, w http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(w).Encode(&res)
}
