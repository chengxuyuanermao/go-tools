package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Println("aabb---")

	tableId := fmt.Sprintf("#%v", RandInterval(10000000, 99999999))
	fmt.Println(tableId)
}

func RandInterval(b1, b2 int32) int32 {
	if b1 == b2 {
		return b1
	}

	min, max := int64(b1), int64(b2)
	if min > max {
		min, max = max, min
	}
	return int32(rand.Int63n(max-min+1) + min)
}
