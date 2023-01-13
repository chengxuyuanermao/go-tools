package main

import (
	"fmt"
	"github.com/chengxuyuanermao/goTools/redBookSpider"
	"strconv"
	"time"
)

func main() {
	redBookSpider.Use()
	fmt.Println(strconv.Itoa(int(time.Now().UnixMilli())))
}
