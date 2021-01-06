// Package assert defines the family of "assert" functions.
package assert

import (
	"fmt"
	"reflect"
	"runtime"
)

func That(condition bool, fmtAndArgs ...interface{}) {
	if !condition {
		panic(newAssertError("condition is not true", fmtAndArgs))
	}
}

func NotNil(i interface{}, fmtAndArgs ...interface{}) {
	if !(i==nil || reflect.ValueOf(i).IsNil()) {
		panic(newAssertError(
			fmt.Sprintf("value %v of type %T is not nil", i, i),
			fmtAndArgs))
	}
}

func NoError(err error, fmtAndArgs ...interface{}) {
	if err != nil {
		panic(newAssertError("error is not nil: "+err.Error(), fmtAndArgs))
	}
}

func errorMsg(defaultMsg string, fmtAndArgs ...interface{}) string {
	if len(fmtAndArgs) == 0 {
		return defaultMsg
	} else {
		return fmt.Sprintf(fmtAndArgs[0].(string), fmtAndArgs[1:]...)
	}
}

func newAssertError(defaultMsg string, fmtAndArgs ...interface{}) *assertError {
	return &assertError{
		method: methodName(),
		msg:    errorMsg(defaultMsg, fmtAndArgs),
	}
}

type assertError struct {
	method string
	msg    string
}

func (e *assertError) Error() string {
	return fmt.Sprintf("assert: assertion failed in %s: %s", e.method, e.msg)
}

// methodName returns the name of the calling method, assumed to be three stack
// frames above.
func methodName() string {
	pc, _, _, _ := runtime.Caller(3)
	f := runtime.FuncForPC(pc)
	if f == nil {
		return "unknown method"
	}
	return f.Name()
}
