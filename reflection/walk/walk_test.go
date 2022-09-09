// Package walk walks.
package walk_test

import (
	"testing"

	"github.com/patrickquigley102/lgwt2/reflection/walk"
	"github.com/stretchr/testify/assert"
)

func TestWalk(t *testing.T) {
	t.Parallel()
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"happy, two string fields",
			args{
				aStruct{"a", "b"},
			},
			[]string{"a", "b"},
		},
		{
			"happy: ignore non-string field",
			args{
				struct {
					A string
					N int
				}{"a", 1},
			},
			[]string{"a"},
		},
		{
			"happy, nested struct",
			args{
				bStruct{
					aStruct{"a", "b"},
					"c",
					"d",
				},
			},
			[]string{"a", "b", "c", "d"},
		},
		{
			"happy, pointers to string and structs",
			args{
				&aStruct{"a", "b"},
			},
			[]string{"a", "b"},
		},
		{"slice", args{[]aStruct{{"a", "b"}}}, []string{"a", "b"}},
		{"array", args{[1]aStruct{{"a", "b"}}}, []string{"a", "b"}},
		{"map", args{map[string]string{"a": "b"}}, []string{"b"}},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := []string{}
			fn := func(i string) {
				got = append(got, i)
			}
			walk.Walk(test.args.x, fn)
			assert.Equal(t, test.want, got)
		})
	}
}

type aStruct struct {
	A string
	B string
}

type bStruct struct {
	AStruct aStruct
	C       string
	D       string
}
