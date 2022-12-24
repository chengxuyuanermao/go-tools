package errors

import (
	"errors"
	"fmt"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		xx := errors.New("math: square root of negative number")
		return 0, errors.Unwrap(xx)
	}
	// 实现
	return 0, nil
}

func Use() {
	//_, err := Sqrt(-1)
	//fmt.Printf("%#v \n", err)
	//fmt.Printf("%v", err)

	//testWrap()
	//testIs()
	testAs()
}

// 嵌套错误处理：https://www.flysnow.org/2019/09/06/go1.13-error-wrapping
// 嵌套错误的好处：给error增加一些附加文本，增加堆栈信息等
/**
输出
Wrap了一个错误：原始错误e
原始错误e
*/
func testWrap() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误：%w", e) // 注意这里的 %w ，生成 Wrapping Error 的关键
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w)) // Unwrap 相当于剥开了一层
}

func testIs() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误：%w", e)

	//res := errors.Is(w, e)                   // 即使嵌套了在内层，也能判断出来，为true
	res := errors.Is(w, errors.New("原始错误e")) // 由于new返回出来的是指针，所以为false

	fmt.Println(res)
}

func testAs() {
	e := newMyError("我定义的错误")
	w := fmt.Errorf("Wrap了一个错误：%w", e)

	//把 w 转为另外一个 myError 类型
	var err *myError
	if res := errors.As(w, &err); res {
		fmt.Println(err) // 输出：我定义的错误
	}
}

// ---  以下是我自己实现的error ----
func newMyError(s string) error {
	return &myError{s}
}

type myError struct {
	msg string
}

func (this *myError) Error() string {
	return this.msg
}
