package main

import (
	"encoding/json"
	"fmt"
)

type cjsonElem struct {
	GuildName string `json:"guildName"`
	NickName  string `json:"nickName"`
	UserId    int32  `json:"userId"`
	UserName  string `json:"userName"`
}

type cjson struct {
	Code    int32       `json:"code"`
	Data    []cjsonElem `json:"data"`
	Message string      `json:"message"`
}

func main() {
	var str string
	str =
		`{
		"code": 10,
		"data": [
			{
				"guildName": "Test1",
				"nickName": "def",
				"userId": 10025,
				"userName": "abc"
			},
			{
				"guildName": "Test2",
				"nickName": "def",
				"userId": 10026,
				"userName": "yxz"
			}
		],
		"message": "success"
	}`
	val := cjson{}
	err := json.Unmarshal([]byte(str), &val)
	if err != nil {
		fmt.Println("Unmarshal failed, err:", err)
		return
	}
	fmt.Println("val:", val)
	fmt.Println(val.Data[0].GuildName)
}

/**
输出结果：
val: {10 [{Test1 def 10025 abc} {Test2 def 10026 yxz}] success}
Test1
*/
