// Package async contains various helpers that allow turning synchronous
// computation into asynchronous or semi-asynchronous computations.
package async

import "context"

// WithContext turns a synchronous producer into a synchronous one that honor
// the context. For example, if the producer requires 10s to finish and context
// hits deadline after 5s, this function will return after 5s with the
// context.DeadlineExceeded error.
func WithContext(ctx context.Context, producer func() (interface{}, error)) (interface{}, error) {
	asyncProducer := func() <-chan resultAndError {
		// Note, that the channel cannot have size=0 as otherwise the goroutine in
		// this function may hang indefinitely (and leak memory) if the context is
		// done before the producer finishes.
		out := make(chan resultAndError, 1)
		go func() {
			result, err := producer()
			out <- resultAndError{result: result, err: err}
		}()
		return out
	}
	select {
	case resErr := <-asyncProducer():
		return resErr.result, resErr.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type resultAndError struct {
	result interface{}
	err    error
}
