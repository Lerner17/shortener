package routes

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/Lerner17/shortener/internal/config"
)

var cgf = config.GetConfig()

func TestRedirectHandler(t *testing.T) {
	type args struct {
		db URLGetter
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RedirectHandler(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedirectHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
