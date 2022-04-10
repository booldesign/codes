package codes

import (
	"fmt"
	"sync"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/4/10 16:10
 * @Desc:
 */

// codeMap 错误码映射 map[int]coder
var codeMap = sync.Map{}

// Register 注册codeMap
// 存在则覆盖
func Register(coder Coder) {
	if coder.Code() == 0 {
		panic("code cannot be zero")
	}
	codeMap.Store(coder.Code(), coder)
}

// MustRegister 注册codes
// 存在则 panic
func MustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic("code cannot be zero")
	}
	if _, loaded := codeMap.LoadOrStore(coder.Code(), coder); loaded {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}
}

// ParseCoder 解析错误 *withCode
// 未知类型错误将被解析为 ErrUnknown
func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}

	if v, ok := err.(*withCode); ok {
		if coder, ok := codeMap.Load(v.code); ok {
			return coder.(Coder)
		}
	}

	return unknownCoder
}

// IsCode 错误链中是否包含给定的错误代码
func IsCode(err error, code int) bool {
	if v, ok := err.(*withCode); ok {
		if v.code == code {
			return true
		}

		if v.cause != nil {
			return IsCode(v.cause, code)
		}

		return false
	}

	return false
}

func init() {
	codeMap.Store(unknownCoder.Code(), unknownCoder)
}
