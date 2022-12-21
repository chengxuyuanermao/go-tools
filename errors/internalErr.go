package errors

import (
	"errors"
	"fmt"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return 0, nil
}

func Use() {
	_, err := Sqrt(-1)
	fmt.Printf("%#v \n", err)
	fmt.Printf("%v", err)
}
