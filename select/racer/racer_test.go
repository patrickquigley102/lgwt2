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
	baseSer := serv(5 * time.Millisecond)
	tests := []struct {
		name      string
		want      string
		timeout   time.Duration
		ser2      *httptest.Server
		assertion assert.ErrorAssertionFunc
	}{
		{
			"success",
			baseSer.URL,
			(200 * time.Millisecond),
			serv(20 * time.Millisecond),
			assert.NoError,
		},
		{
			"timeout failures",
			"",
			(1 * time.Millisecond),
			serv(10 * time.Millisecond),
			assert.Error,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := racer.Racer(baseSer.URL, test.ser2.URL, test.timeout)
			assert.Equal(t, test.want, got)
			test.assertion(t, err)

			baseSer.Close()
			test.ser2.Close()
		})
	}
}

func serv(wait time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(wait)
			},
		),
	)
}
