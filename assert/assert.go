// Package assert defines the family of "assert" functions.
package assert

import "fmt"

func Assert(condition bool) {
	if !condition {
		panic("Condition is not true")
	}
}

func Must(err error) {
	if err != nil {
		panic(fmt.Sprintf("Error is not nil: %v", err))
	}
}
