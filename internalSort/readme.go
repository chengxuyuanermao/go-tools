package internalSort

/**
Go语言的sort包提供了对切片和用户自定义集合进行排序的函数。该包实现了常见的排序算法，如快速排序和堆排序，并提供了用于自定义排序行为的接口。

以下是sort包中最常用的函数和接口：

func Ints(a []int)：对int类型的切片进行升序排序。
func Float64s(a []float64)：对float64类型的切片进行升序排序。
func Strings(a []string)：对string类型的切片进行升序排序。
func IntsAreSorted(a []int) bool：检查int类型的切片是否已经按升序排序。
func Float64sAreSorted(a []float64) bool：检查float64类型的切片是否已经按升序排序。
func StringsAreSorted(a []string) bool：检查string类型的切片是否已经按升序排序。
在这些函数之上，sort包还提供了一些用于自定义排序行为的接口，以便对用户定义的类型进行排序。其中最常用的接口是sort.Interface接口，该接口定义了三个方法：

Len() int：返回集合中的元素个数。
Less(i, j int) bool：报告索引为i的元素是否应该排在索引为j的元素之前。
Swap(i, j int)：交换索引为i和j的元素。
用户可以实现sort.Interface接口以定义自己的排序逻辑，然后使用sort.Sort函数对切片进行排序。

*/
