
/*
errors 包使用说明
它的使用非常简单，如果我们要新生成一个错误，可以使用New函数,生成的错误，自带调用堆栈信息。
func New(message string) error

如果有一个现成的error，我们需要对他进行再次包装处理，这时候有三个函数可以选择。
//只附加新的信息
func WithMessage(err error, message string) error
//只附加调用堆栈信息
func WithStack(err error) error
//同时附加堆栈和信息
func Wrap(err error, message string) error


其实上面的包装，很类似于Java的异常包装，被包装的error，其实就是Cause,在前面的章节提到错误的根本原因，就是这个Cause。所以这个错误处理库为我们提供了Cause函数让我们可以获得最根本的错误原因。
func Cause(err error) error {
type causer interface {
Cause() error
}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}


使用for循环一直找到最根本（最底层）的那个error。
以上的错误我们都包装好了，也收集好了，那么怎么把他们里面存储的堆栈、错误原因等这些信息打印出来呢？其实，这个错误处理库的错误类型，都实现了Formatter接口，我们可以通过fmt.Printf函数输出对应的错误信息。

%s,%v //功能一样，输出错误信息，不包含堆栈 // 示例：2022/04/09 20:18:20 unable to serve HTTP POST request for customer 1, unable to insert customer contract 1: unable to commit transaction

%q //输出的错误信息"带引号"，不包含堆栈 // 示例：2022/04/09 20:18:03 unable to serve HTTP POST request for customer 1, "unable to insert customer contract 1: unable to commit transaction"

%+v //输出错误信息和堆栈
// 示例：2022/04/09 20:16:15 unable to commit transaction
//github.com/booldesign/errors/example.dbQuery
//        /Users/booldesign/Applications/gowork/src/github.com/booldesign/errors/example/error_test.go:127
//github.com/booldesign/errors/example.insert
//        /Users/booldesign/Applications/gowork/src/github.com/booldesign/errors/example/error_test.go:117
//github.com/booldesign/errors/example.postHandler

以上如果有循环包装错误类型的话，会递归的把这些错误都会输出。

*/