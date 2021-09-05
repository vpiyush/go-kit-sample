package inmem

import (
	"reflect"
	"sync"
	"testing"
)

func Test_storage_Insert(t *testing.T) {
	type fields struct {
		items map[string]string
		mu    sync.RWMutex
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Pair
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &storage{
				items: tt.fields.items,
				mu:    tt.fields.mu,
			}
			got, err := p.Insert(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("storage.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storage.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
