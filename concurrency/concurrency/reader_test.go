package concurrency_test

import (
	"log"
	"os"
	"testing"

	"github.com/patrickquigley102/lgwt2/concurrency/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestFileToArray(t *testing.T) {
	t.Parallel()
	type args struct {
		fileName string
	}
	tests := []struct {
		name      string
		args      args
		want      []string
		assertion assert.ErrorAssertionFunc
	}{
		{"success", args{file(t, "a,b,c")}, []string{"b"}, assert.NoError},
		{"no second colunm", args{file(t, "a")}, nil, assert.Error},
		{"invalid csv", args{file(t, "a\nb\n")}, nil, assert.Error},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := concurrency.FileToArray(test.args.fileName)
			test.assertion(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func file(t *testing.T, content string) string {
	t.Helper()
	file, err := os.CreateTemp("", "urls*.csv")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	return file.Name()
}
