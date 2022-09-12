// Package server serves.
package server

import (
	"fmt"
	"net/http"
)

// Server serves data from Store.
func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		dataChan := make(chan string, 1)

		go func() {
			dataChan <- store.Fetch()
		}()

		select {
		case data := <-dataChan:
			fmt.Fprint(writer, data)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

// Store is an interface.
type Store interface {
	Fetch() string
	Cancel()
}
