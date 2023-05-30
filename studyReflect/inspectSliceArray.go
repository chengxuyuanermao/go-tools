package studyReflect

import (
	"fmt"
	"reflect"
)

/**
透视切片或数组组成，需要以下方法：

reflect.Value.Len()：返回数组或切片的长度；
reflect.Value.Index(i)：返回第i个元素的reflect.Value值；
然后对这个reflect.Value判断Kind()进行处理。
示例：
同样地Len()和Index(i)方法只能在原对象是切片，数组或字符串时才能调用，其他类型会panic。
*/

func Use6() {
	inspectSliceArray([]int{1, 2, 3})
}

func inspectSliceArray(m interface{}) {
	v := reflect.ValueOf(m)
	v.Type()
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("元素有：%v \n", elem.Interface())
	}
}
