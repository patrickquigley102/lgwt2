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
		name      string
		ser1      *httptest.Server
		ser2      *httptest.Server
		assertion assert.ErrorAssertionFunc
	}{
		{"test", testServer(0), testServer(20 * time.Millisecond), assert.NoError},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := racer.Racer(test.ser1.URL, test.ser2.URL)
			assert.Equal(t, test.ser1.URL, got)
			test.assertion(t, err)

			test.ser1.Close()
			test.ser2.Close()
		})
	}
}

func testServer(wait time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(wait)
			},
		),
	)
}
