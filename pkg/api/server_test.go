package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewHttpServer(t *testing.T) {
	srv := httptest.NewServer(NewHttpServer())
	defer srv.Close()
	tests := []struct {
		name, method, url, body, want string
	}{
		{"InsertSuccess", "POST", srv.URL + "/pair", `{"key":"key-1","value":"value-1"}`, `{"key":"key-1","value":"value-1"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//	rr := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			resp, _ := http.DefaultClient.Do(req)
			body, _ := ioutil.ReadAll(resp.Body)
			if want, have := tt.want, strings.TrimSpace(string(body)); want != have {
				t.Errorf("NewHttpServer() = %v, want %v", have, tt.want)
			}
		})
	}
}
