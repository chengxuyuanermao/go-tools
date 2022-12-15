package pinyin

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func Use() {
	hans := "中国人"
	a := pinyin.NewArgs() // a相当于参数配置

	// 默认
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren]]

	// 包含声调
	a.Style = pinyin.Tone // 配置有声调
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhōng] [guó] [rén]]

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng] [guo2] [re2n]]
}
