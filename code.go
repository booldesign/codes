package codes

import (
	"net/http"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/4/10 11:03
 * @Desc:
 */

// Coder code 接口
type Coder interface {
	// HTTPStatus HTTP status
	HTTPStatus() int

	// Code err code
	Code() int

	// String err message
	String() string
}

var unknownCoder = defaultCoder{
	HTTP: http.StatusInternalServerError,
	C:    1,
	Msg:  "服务内部错误，请稍后再试...",
}

type defaultCoder struct {
	// HTTP status that should be used for the associated error code.
	HTTP int

	// C refers to the integer code of the ErrCode.
	C int

	// error text
	Msg string
}

// HTTPStatus 返回 http status
func (coder defaultCoder) HTTPStatus() int {
	if coder.HTTP == 0 {
		return http.StatusInternalServerError
	}
	return coder.HTTP
}

// Code 返回 code 值
func (coder defaultCoder) Code() int {
	return coder.C
}

// String 返回 err message
func (coder defaultCoder) String() string {
	return coder.Msg
}
