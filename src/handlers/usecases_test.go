package handlers

import (
	"net/http"
	"testing"
)

func TestShowAllEmployees(t *testing.T) {
	var (
		w http.ResponseWriter
		r *http.Request
	)

	ShowAllEmployees(w, r)
}
