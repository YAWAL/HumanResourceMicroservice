package handlers

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/YAWAL/HumanResourceMicroservice/src/repository"
)

func TestShowAllEmployeesByPosition(t *testing.T) {
	type args struct {
		er repository.EmployeeRepository
	}
	tests := []struct {
		name string
		args args
		want func(http.ResponseWriter, *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShowAllEmployeesByPosition(tt.args.er); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowAllEmployeesByPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
