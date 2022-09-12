// Package server serves.
package server

import (
	"fmt"
	"net/http"
)

// Server serves data from Store.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}

// Store is an interface.
type Store interface {
	Fetch() string
}
