// Package concurrency is a package.
package concurrency_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/patrickquigley102/lgwt2/concurrency/concurrency"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckWebsites(t *testing.T) {
	t.Parallel()
	type args struct {
		wc   *mockWebsiteChecker
		urls []string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		{
			"test",
			args{new(mockWebsiteChecker), []string{"a", "b"}},
			map[string]bool{"a": true, "b": false},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			test.args.wc.On("Check", "a").Return(true)
			test.args.wc.On("Check", "b").Return(false)

			assert.Equal(
				t,
				test.want,
				concurrency.CheckWebsites(test.args.wc, test.args.urls),
			)
			test.args.wc.AssertExpectations(t)
		})
	}
}

type mockWebsiteChecker struct {
	mock.Mock
}

func (m *mockWebsiteChecker) Check(url string) bool {
	args := m.Called(url)

	return args.Bool(0)
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(new(slowWebsiteChecker), urls)
	}
}

type slowWebsiteChecker struct{}

func (swc slowWebsiteChecker) Check(string) bool {
	time.Sleep(20 * time.Millisecond)

	return true
}

func TestWebsiteCheck_Check(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		want           bool
		testServerFunc testServerFunc
	}{
		{"200", true, test200Server},
		{"200", false, testNot200Server},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ts := test.testServerFunc(t)
			defer ts.Close()
			t.Parallel()
			wc := concurrency.WebsiteCheck{}
			assert.Equal(t, test.want, wc.Check(ts.URL))
		})
	}
}

type testServerFunc func(t *testing.T) *httptest.Server

func test200Server(t *testing.T) *httptest.Server {
	t.Helper()

	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {},
		),
	)
}

func testNot200Server(t *testing.T) *httptest.Server {
	t.Helper()

	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
			},
		),
	)
}
