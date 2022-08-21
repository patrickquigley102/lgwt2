// Package structs for learnings
package structs

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangle_Perimeter(t *testing.T) {
	t.Parallel()
	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"zero", fields{0.0, 0.0}, 0.0},
		{"rect", fields{1.0, 1.0}, 4.0},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			r := Rectangle{
				Width:  test.fields.Width,
				Height: test.fields.Height,
			}
			assert.Equal(t, test.want, r.Perimeter())
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	t.Parallel()
	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"zero", fields{0.0, 0.0}, 0.0},
		{"rect", fields{1.0, 1.0}, 1.0},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			r := Rectangle{
				Width:  test.fields.Width,
				Height: test.fields.Height,
			}
			assert.Equal(t, test.want, r.Area())
		})
	}
}

func TestCircle_Area(t *testing.T) {
	t.Parallel()
	type fields struct {
		Radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"zero", fields{0.0}, 0.0},
		{"rect", fields{1.0}, math.Pi},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := Circle{
				Radius: test.fields.Radius,
			}
			assert.Equal(t, test.want, c.Area())
		})
	}
}
