package concurrency_test

import (
	"log"
	"os"
	"testing"

	"github.com/patrickquigley102/lgwt2/concurrency/concurrency"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDownloadCSV(t *testing.T) {
	var filePathName string = "/src/personal/go/lgwt2/" +
		"concurrency/concurrency/urls.csv"

	t.Parallel()
	type args struct {
		fs *mockFS
	}
	tests := []struct {
		name      string
		args      args
		want      *os.File
		assertion assert.ErrorAssertionFunc
	}{
		{"create blank file", args{new(mockFS)}, new(os.File), assert.NoError},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			test.args.fs.On("UserHomeDir").Return("", nil)
			test.args.fs.On("Create", filePathName).Return(test.want, nil)

			got, err := concurrency.DownloadCSV(test.args.fs)

			test.assertion(t, err)
			assert.Equal(t, test.want, got)
			test.args.fs.AssertExpectations(t)
		})
	}
}

type mockFS struct {
	mock.Mock
}

func (m *mockFS) Create(name string) (*os.File, error) {
	args := m.Called(name)
	file, err := args.Get(0).(*os.File)
	if !err {
		log.Fatal(err)
	}

	return file, args.Error(1)
}

func (m *mockFS) UserHomeDir() (string, error) {
	args := m.Called()

	return args.String(0), args.Error(1)
}
