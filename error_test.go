package codes

import (
	"github.com/pkg/errors"
	"io"
	"reflect"
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/4/10 15:28
 * @Desc:
 */

var ErrRecordNotFound = 404

func TestWithCode(t *testing.T) {
	tests := []struct {
		code     int
		message  string
		wantType string
		wantCode int
	}{
		{ErrRecordNotFound, "record not found", "*withCode", ErrRecordNotFound},
	}

	for _, tt := range tests {
		got := WithCode(tt.code, tt.message)
		err, ok := got.(*withCode)
		if !ok {
			t.Errorf("WithCode(%v, %q): error type got: %T, want %s", tt.code, tt.message, got, tt.wantType)
		}

		if err.code != tt.wantCode {
			t.Errorf("WithCode(%v, %q): got: %v, want %v", tt.code, tt.message, err.code, tt.wantCode)
		}
	}
}

func TestWithCodef(t *testing.T) {
	tests := []struct {
		code       int
		format     string
		args       string
		wantType   string
		wantCode   int
		wangString string
	}{
		{ErrRecordNotFound, "id: %s, record not found", "1", "*withCode", ErrRecordNotFound, `id: 1, record not found`},
	}

	for _, tt := range tests {
		got := WithCode(tt.code, tt.format, tt.args)
		err, ok := got.(*withCode)
		if !ok {
			t.Errorf("WithCode(%v, %q %q): error type got: %T, want %s", tt.code, tt.format, tt.args, got, tt.wantType)
		}

		if err.code != tt.wantCode {
			t.Errorf("WithCode(%v, %q %q): got: %v, want %v", tt.code, tt.format, tt.args, err.code, tt.wantCode)
		}

		if got.Error() != tt.wangString {
			t.Errorf("WithCode(%v, %q %q): got: %v, want %v", tt.code, tt.format, tt.args, got.Error(), tt.wangString)
		}
	}
}

type nilError struct{}

func (nilError) Error() string { return "nil error" }

func TestCause(t *testing.T) {
	x := errors.New("error")
	tests := []struct {
		err  error
		want error
	}{
		{
			// nil error is nil
			err:  nil,
			want: nil,
		}, {
			// explicit nil error is nil
			err:  (error)(nil),
			want: nil,
		}, {
			// typed nil is nil
			err:  (*nilError)(nil),
			want: (*nilError)(nil),
		}, {
			// uncaused error is unaffected
			err:  io.EOF,
			want: io.EOF,
		}, {
			// caused error returns cause
			err:  errors.Wrap(io.EOF, "ignored"),
			want: io.EOF,
		}, {
			err:  x, // return from errors.New
			want: x,
		}, {
			errors.WithMessage(nil, "whoops"),
			nil,
		}, {
			errors.WithMessage(io.EOF, "whoops"),
			io.EOF,
		}, {
			errors.WithStack(nil),
			nil,
		}, {
			errors.WithStack(io.EOF),
			io.EOF,
		}, {
			(io.EOF),
			io.EOF,
		},
	}

	for i, tt := range tests {
		got := Cause(tt.err)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("test %d: got %#v, want %#v", i+1, got, tt.want)
		}
	}
}
