# Package codes

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/booldesign/gvalid)
![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


本包的目的是解决日常api架构的code/message格式，部分借鉴 `github.com/pkg/errors` 包

```
// 创建一个带 `code` 和 `message` 的 error
func WithCode(code int, format string, args ...interface{}) error

// 同时附加堆栈和code、信息
func WrapCode(err error, code int, format string, args ...interface{}) error
```

## 使用示例

 ```
// 注册实现 coder 接口的 Err
MustRegister(coder)

// 解析 coder
ParseCoder(err error)
 ```
