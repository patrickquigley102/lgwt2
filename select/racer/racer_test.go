package racer_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/patrickquigley102/lgwt2/select/racer"
	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		server1 testServerFunc
		server2 testServerFunc
	}{
		{"test", testServer, testSlowServer},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ts1 := test.server1(t)
			ts2 := test.server2(t)

			assert.Equal(t, ts1.URL, racer.Racer(ts1.URL, ts2.URL))

			ts1.Close()
			ts2.Close()
		})
	}
}

type testServerFunc func(t *testing.T) *httptest.Server

func testServer(t *testing.T) *httptest.Server {
	t.Helper()

	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {},
		),
	)
}

func testSlowServer(t *testing.T) *httptest.Server {
	t.Helper()

	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(20 * time.Millisecond)
			},
		),
	)
}
