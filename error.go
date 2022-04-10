package codes

import (
	"fmt"
	"io"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/4/9 20:42
 * @Desc:
 */

type withCode struct {
	err   error
	code  int
	cause error
	*stack
}

func WithCode(code int, format string, args ...interface{}) error {
	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		stack: callers(),
	}
}

// WrapCode 同时附加堆栈和信息
func WrapCode(err error, code int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		cause: err,
		stack: callers(),
	}
}

// Error return the externally-safe error message.
func (w *withCode) Error() string {
	return w.err.Error()
	/*if w.cause == nil {
		return w.err.Error()
	}
	return Cause(w.cause).Error()*/
	//return fmt.Sprintf("%v", w)
}

// Cause return the cause of the withCode error.
func (w *withCode) Cause() error {
	return w.cause
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withCode) Unwrap() error {
	return w.cause
}

// Format print 时调用
func (w *withCode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			io.WriteString(s, w.err.Error())
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error())
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}

		if cause.Cause() == nil {
			break
		}

		err = cause.Cause()
	}
	return err
}
