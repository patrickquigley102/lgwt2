package server_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/patrickquigley102/lgwt2/context/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const millisecondsToWaitBeforeCancel = 5 * time.Millisecond

func TestServer(t *testing.T) {
	t.Parallel()
	type args struct {
		store *mockStore
	}
	tests := []struct {
		name      string
		args      args
		want      string
		cancelled bool
	}{
		{"success, fast", args{newMockStore(0)}, "test", false},
		{"success, slow and cancelled", args{newMockStore(100)}, "", true},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			svr := server.Server(test.args.store)

			test.args.store.On("Fetch").Return(test.want)
			if test.cancelled {
				test.args.store.On("Cancel")
			}

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			context, cancelFunc := context.WithCancel(request.Context())
			time.AfterFunc(millisecondsToWaitBeforeCancel, cancelFunc)
			request = request.WithContext(context)

			response := httptest.NewRecorder()

			svr.ServeHTTP(response, request)

			assert.Equal(t, test.want, response.Body.String())
			test.args.store.AssertExpectations(t)
		})
	}
}

func newMockStore(sleepLength int) *mockStore {
	return &mockStore{sleepLength: time.Duration(sleepLength) * time.Millisecond}
}

type mockStore struct {
	mock.Mock
	sleepLength time.Duration
}

func (m *mockStore) Fetch() string {
	args := m.Called()

	time.Sleep(m.sleepLength)

	return args.String(0)
}

func (m *mockStore) Cancel() {
	m.Called()
}
