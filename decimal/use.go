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
