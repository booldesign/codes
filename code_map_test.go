package codes

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/4/10 16:13
 * @Desc:
 */

func TestRegister(t *testing.T) {
	type args struct {
		coder Coder
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test Register", args: args{coder: defaultCoder{HTTP: http.StatusInternalServerError, C: 2, Msg: "服务内部错误，请稍后再试..."}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.coder)
			if v, ok := codeMap.Load(tt.args.coder.Code()); !ok {
				t.Fatal("register() err")
			} else if !reflect.DeepEqual(tt.args.coder, v) {
				t.Fatal("register() err")
			}
		})
	}
}

func TestMustRegister(t *testing.T) {
	type args struct {
		coder Coder
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test MustRegister", args: args{coder: defaultCoder{HTTP: http.StatusInternalServerError, C: 2, Msg: "服务内部错误，请稍后再试..."}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.coder)
			defer func() {
				if err := recover(); err != nil {
					t.Skipped()
				} else {
					t.Fatalf("MustRegister() error:%v", err)
				}
			}()
			MustRegister(tt.args.coder)
		})
	}
}

func TestParseCoder(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		wantHTTPCode int
		wantString   string
		wantCode     int
	}{
		{"test ParseCoder", fmt.Errorf("db error"), 500, "服务内部错误，请稍后再试...", 1},
		{"test ParseCoder", WithCode(unknownCoder.Code(), "internal error message"), 500, "服务内部错误，请稍后再试...", 1},
		{"test ParseCoder", WrapCode(fmt.Errorf("db error"), unknownCoder.Code(), "internal error message"), 500, "服务内部错误，请稍后再试...", 1},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coder := ParseCoder(tt.err)
			if coder.HTTPStatus() != tt.wantHTTPCode {
				t.Errorf("TestCodeParse(%d): got %q, want: %q", i, coder.HTTPStatus(), tt.wantHTTPCode)
			}

			if coder.String() != tt.wantString {
				t.Errorf("TestCodeParse(%d): got %q, want: %q", i, coder.String(), tt.wantString)
			}

			if coder.Code() != tt.wantCode {
				t.Errorf("TestCodeParse(%d): got %q, want: %q", i, coder.Code(), tt.wantCode)
			}
		})
	}
}

func TestIsCode(t *testing.T) {
	type args struct {
		err  error
		code int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test IsCode", args: args{err: WrapCode(WithCode(1, "服务内部错误，请稍后再试..."), 2, "嵌套err"), code: 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCode(tt.args.err, tt.args.code); got != tt.want {
				t.Errorf("IsCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
