package projectEncrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// 向平台方发起提现请求
type platformWithdrawRequest struct {
	id        string
	order     string
	amount    string
	notifyurl string
	notes     string
	accname   string
	accno     string
	bankcode  string
	sign      string
}

func TestGetSign() {
	p := &platformWithdrawRequest{
		id:        "50053374",
		order:     "20230711013026933017",
		amount:    "100.00",
		notifyurl: "https://bubble.paradisebubble.site/order/cashback_v2",
		notes:     `{"email":"520dfdf155@gmail.com","kind":"upi","mobile":"9938454253"}`,
		accname:   "Sunil ramesh dodwa",
		accno:     "20297405373",
		bankcode:  "upi",
		sign:      "",
	}
	res := getSign(p)
	fmt.Println(res)
}

func getSign(req *platformWithdrawRequest) string {
	resReqMap := make(map[string]string)
	resReqMap["id"] = req.id
	resReqMap["order"] = req.order
	resReqMap["amount"] = req.amount
	resReqMap["notifyurl"] = req.notifyurl
	resReqMap["notes"] = req.notes
	resReqMap["accname"] = req.accname
	resReqMap["accno"] = req.accno
	resReqMap["bankcode"] = req.bankcode
	keys := make([]string, 0)
	for k, _ := range resReqMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sortedData := make([]string, 0)
	for _, key := range keys {
		sortedData = append(sortedData, key+"="+resReqMap[key])
	}
	sortedData = append(sortedData, "key=D3QXUUVWPFO5GNJEISC0JTTTLQDZKMMC") // 末尾加上key
	reqStr := strings.Join(sortedData, "&")                                 // 串联
	hash := md5.Sum([]byte(reqStr))
	sign := hex.EncodeToString(hash[:])
	return sign
}
