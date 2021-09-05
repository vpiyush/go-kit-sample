package services

import (
	"reflect"
	"testing"

	"github.com/vpiyush/go-kit-sample/pkg/storage/inmem"
)

func Test_service_Insert(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *inmem.Pair
		wantErr bool
	}{
		{"t1", fields{inmem.NewStorage()}, args{"key-1", "value-1"}, &inmem.Pair{Key: "key-1", Value: "value-1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			got, err := s.Insert(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Get(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *inmem.Pair
		wantErr bool
	}{
		{"GetFailed", fields{inmem.NewStorage()}, args{"key-1"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
