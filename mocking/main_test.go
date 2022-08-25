package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCountdown(t *testing.T) {
	t.Parallel()
	type args struct {
		s *mockSleeper
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{"test", args{new(mockSleeper)}, "3\n2\n1\nGo!"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			test.args.s.On("Sleep").Times(startCount)

			Countdown(w, test.args.s)

			assert.Equal(t, test.wantW, w.String())
			test.args.s.AssertExpectations(t)
		})
	}
}

type mockSleeper struct {
	mock.Mock
}

func (m *mockSleeper) Sleep() {
	m.Called()
}

func TestConfigurableSleeper_Sleep(t *testing.T) {
	t.Parallel()
	mockSleep := new(mockConfigurableSleeper)
	duration := time.Second
	mockSleepFunc := mockSleep.mockSleepFunc
	configurableSleeper := ConfigurableSleeper{
		duration: time.Second,
		sleep:    mockSleepFunc,
	}

	mockSleep.On("mockSleepFunc", duration)
	configurableSleeper.Sleep()
	mockSleep.AssertExpectations(t)
}

type mockConfigurableSleeper struct {
	mock.Mock
}

func (m *mockConfigurableSleeper) mockSleepFunc(duration time.Duration) {
	m.Called(duration)
}
