package example

import (
	"net/http"

	"github.com/booldesign/codes"
)

// ErrCode implements `github.com/booldesign/codes`.Coder interface.
type ErrCode struct {
	// HTTP status
	HTTP int

	// C err code
	C int

	// err text
	Msg string
}

var _ codes.Coder = &ErrCode{}

// HTTPStatus 返回 http status
func (coder ErrCode) HTTPStatus() int {
	if coder.HTTP == 0 {
		return http.StatusInternalServerError
	}
	return coder.HTTP
}

// Code 返回 code 值
func (coder ErrCode) Code() int {
	return coder.C
}

// String 返回 err message
func (coder ErrCode) String() string {
	return coder.Msg
}

// 注册 coder
func register(code int, httpStatus int, message string) {
	if found := inSlice([]int{http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden,
		http.StatusNotFound, http.StatusInternalServerError}, httpStatus); !found {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}
	coder := &ErrCode{
		C:    code,
		HTTP: httpStatus,
		Msg:  message,
	}

	codes.MustRegister(coder)
}

func inSlice(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}
