package async_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/mm4tt/goutil/async"
)

func TestWithContext(t *testing.T) {
	var producerError = errors.New("producer: error")

	testcases := []struct {
		name         string
		producer     func() (interface{}, error)
		ctxModifier  func(context.Context) context.Context
		wantedResult interface{}
		wantedError  error
	}{
		{
			name: "producer returns result",
			producer: func() (interface{}, error) {
				return "result", nil
			},
			wantedResult: "result",
		},
		{
			name: "producer returns error",
			producer: func() (interface{}, error) {
				return nil, producerError
			},
			wantedError: producerError,
		},
		{
			name: "context hits deadline before producer finishes",
			producer: func() (interface{}, error) {
				time.Sleep(10 * time.Millisecond)
				return "result", nil
			},
			ctxModifier: func(ctx context.Context) context.Context {
				ctx, _ = context.WithTimeout(ctx, 5*time.Millisecond)
				return ctx
			},
			wantedError: context.DeadlineExceeded,
		},
		{
			name: "context cancelled",
			producer: func() (interface{}, error) {
				return "result", nil
			},
			ctxModifier: func(ctx context.Context) context.Context {
				ctx, cancel := context.WithCancel(ctx)
				cancel()
				return ctx
			},
			wantedError: context.Canceled,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			if tc.ctxModifier != nil {
				ctx = tc.ctxModifier(ctx)
			}
			gotResult, gotError := async.WithContext(ctx, tc.producer)
			if gotResult != tc.wantedResult {
				t.Errorf("wanted result: '%v', got result: '%v'", tc.wantedResult, gotResult)
			}
			if gotError != tc.wantedError {
				t.Errorf("wanted error: '%v', got error: '%v'", tc.wantedError, gotError)
			}
		})
	}
}
