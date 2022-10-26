package decimal

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// Use https://github.com/shopspring/decimal
func Use() {
	p1 := decimal.NewFromFloat(1.13)
	p2 := float64(1.13)
	fmt.Printf("有问题，没使用Decimal：%v； 没问题使用了：%v", p2*100, p1.Mul(decimal.NewFromInt(100)))
}

func Main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	num := decimal.NewFromInt(3)

	fmt.Println("0.1+0.3=", 0.1+float64(0.2))
	fmt.Println("fix: 0.1+0.3=", decimal.NewFromFloat(0.1).Add(decimal.NewFromFloat(0.2)))

	fmt.Println("add:", price.Add(num)) // 加
	fmt.Println("Sub:", price.Sub(num)) // 减
	fmt.Println("mul:", price.Mul(num)) // 乘法，price*quantity = 136.02 * 3 = 408.06
	fmt.Println("div:", price.Div(num)) // 除

	fmt.Println("3 div 1.3:", num.Div(decimal.NewFromFloat(1.3))) // 除不尽头

	/**
	add: 139.02
	div: 45.34
	mul: 408.06
	Sub: 133.02
	*/
}
