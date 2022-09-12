package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patrickquigley102/lgwt2/context/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServer(t *testing.T) {
	t.Parallel()
	type args struct {
		store *mockStore
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{new(mockStore)}, "test"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			svr := server.Server(test.args.store)

			test.args.store.On("Fetch").Return(test.want)

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			response := httptest.NewRecorder()

			svr.ServeHTTP(response, request)

			got := response.Body.String()
			assert.Equal(t, test.want, got)
		})
	}
}

type mockStore struct {
	mock.Mock
}

func (m *mockStore) Fetch() string {
	args := m.Called()

	return args.String(0)
}
