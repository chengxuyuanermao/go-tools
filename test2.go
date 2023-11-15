package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
)

type WithdrawV3Request struct {
	SendId      int    `json:"send_id"`
	PriceType   string `json:"price_type"`             // 提现价格类型：cash或coin
	Channel     string `json:"channel"`                // 当前只有pix
	Amount      int    `json:"amount,omitempty"`       // 提现的price
	To          string `json:"to"`                     // 收款账号，可以是： 个⼈税务号，公司税务号，⼿机号，email
	AccountName string `json:"account_name,omitempty"` // 收款⽤户姓名
}

func test() {
	req := &WithdrawV3Request{
		SendId:      1,
		PriceType:   "cash",
		Channel:     "pix",
		Amount:      200,
		To:          "0123456",
		AccountName: "test",
	}
	pureAmount := 200

	// 区分渠道信息
	other := new(otherInfo)
	err := setOtherInfo(req.To, other)
	if err != nil {
		fmt.Printf("withdraw setOtherInfo err:%v", err)
		return
	}

	// ------ 开始操作 ------
	//  请求三方平台 -- 目前只支持巴西
	// 创建内部订单号
	orderId := GenOrderId()
	reqPlatform := &WithdrawRequestV3{
		Pkg:             "com.testslot.bigwin",
		WOrderId:        orderId,
		Uid:             "123456789",
		CallbackUrl:     "https://gs.goldentest.games/withdraw_v3/callback",
		Amount:          fmt.Sprintf("%d.%02d", pureAmount/100, pureAmount%100), //将单位分，换成元
		Mid:             "slotearn",
		CardType:        "pagbank",
		WithdrawChannel: "pix",
		Unit:            "BRL",
		Location:        "BR",

		Ifsc:             other.ifsc,
		PayeeEmail:       other.payeeEmail,
		PayeePhoneNumber: other.payeePhoneNumber,
		BankAccount:      other.bankAccount,
		AccountName:      req.AccountName,
	}
	if err := ReqWithdrawPlatform(reqPlatform); err != nil {
		fmt.Printf("ReqWithdrawPlatform err:%v", err)
		return
	}

	fmt.Println("success")
}

type WithdrawV3Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

const (
	flagReq      = "req"
	flagCallback = "callback"
)

func ReqWithdrawPlatform(req *WithdrawRequestV3) error {
	req.Checksum = getChecksum(*req, flagReq)
	//fmt.Printf("req struct: %+v \n", req)
	resp := httpPost(req)
	if resp == nil {
		return errors.New("httpPost err")
	} else if resp.Status != 1 {
		/**
		返回结果状态,1成功,⾮1失败
		成功时值为success,失败为错误结果
		*/return errors.New(resp.Msg)
	}
	return nil
}

func getChecksum(req interface{}, flagStr string) string {
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	keys := make([]string, 0)
	jsonKeys := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if value.String() == "" {
			continue
		}

		// 获取json中的key
		multiVal := field.Tag.Get("json")
		multiValArr := strings.Split(multiVal, ",")
		jsonKey := strings.Trim(multiValArr[0], "")
		jsonKeys = append(jsonKeys, jsonKey)

		// 获取原本的key
		key := field.Name
		keys = append(keys, key)
	}

	signArr := make([]string, 0)
	sort.Strings(keys)
	sort.Strings(jsonKeys)
	//log.Infof("checkSum sorted keys: %v", keys)
	for idx, key := range keys {
		//fmt.Println("xx--", jsonKeys[idx], key)
		isNeed := true
		if flagStr == flagReq {
			isNeed = isCreateOrderSecretField(jsonKeys[idx])

		} else if flagStr == flagCallback {
			isNeed = isCallbackSecretField(jsonKeys[idx])
		}
		if !isNeed {
			continue
		}

		value := v.FieldByName(key)
		if !value.IsValid() {
			log.Errorf("getChecksum err key:%v", key)
			continue
		}
		signArr = append(signArr, value.String())
	}
	secret := "08ff8shXd21xj0tj2kdeld"
	signArr = append(signArr, secret)
	signStr := strings.Join(signArr, "") // 串联
	//log.Infof("checkSum sorted res: %v, %v", signStr, signArr)
	hash := sha256.Sum256([]byte(signStr))
	checksum := hex.EncodeToString(hash[:])
	return strings.ToLower(checksum)
}

func isCreateOrderSecretField(field string) bool {
	var checkSumField = []string{"w_order_id", "uid", "amount", "unit", "client_type", "payee_email", "payee_phone_number", "callback_url", "mid", "withdraw_channel", "channel_id", "bank_account", "ifsc", "account_name", "upi_id", "card_no", "source"}
	for _, v := range checkSumField {
		if v == field {
			return true
		}
	}
	return false
}

func isCallbackSecretField(field string) bool {
	var notCheckFields = []string{"checksum", "sign"} // 这些除外
	for _, v := range notCheckFields {
		if v == field {
			return false
		}
	}
	return true
}

func httpPost(req *WithdrawRequestV3) *WithdrawV3Response {
	reqUrl := "https://api.goldentest.games/api/ga-api/withdraw/create_order"
	jsonData, _ := json.Marshal(req)
	log.Infof("req jsonData: %v", string(jsonData))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Errorf("WithdrawRequestV3 err:%v", err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Infof("resp body: %v", string(body))

	withdrawResponse := new(WithdrawV3Response)
	err = json.Unmarshal(body, withdrawResponse)
	if err != nil {
		log.Errorf("Unmarshal body err: %v", err)
		return nil
	}
	return withdrawResponse
}

type otherInfo struct {
	ifsc             string
	payeeEmail       string
	payeePhoneNumber string
	bankAccount      string
}

func setOtherInfo(input string, info *otherInfo) error {
	// CPF正则表达式
	cpfRegex := `^\d{3}\.\d{3}\.\d{3}-\d{2}$`

	// CNPJ正则表达式
	cnpjRegex := `^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`

	// 邮箱正则表达式
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 巴西手机号正则表达式
	phoneRegex := `^\d{6,}$`

	// 使用正则表达式进行匹配
	if regexp.MustCompile(cpfRegex).MatchString(input) {
		info.ifsc = "cpf"
		info.bankAccount = input
	} else if regexp.MustCompile(cnpjRegex).MatchString(input) {
		info.ifsc = "cnpj"
		info.bankAccount = input
	} else if regexp.MustCompile(emailRegex).MatchString(input) {
		info.ifsc = "email"
		info.bankAccount = input
		info.payeeEmail = input
	} else if regexp.MustCompile(phoneRegex).MatchString(input) {
		info.ifsc = "phone"
		info.bankAccount = input
		info.payeePhoneNumber = input
	}

	if info.ifsc != "" {
		return nil
	}
	return errors.New("no match type")
}

func GenOrderId() string {
	now := time.Now()
	w := new(strings.Builder)
	w.WriteString(now.Format("20060102150405"))
	fmt.Fprintf(w, "%06d", rand.Intn(1000000))
	return w.String()
}

// 向平台请求
type WithdrawRequestV3 struct {
	Pkg              string `json:"pkg,omitempty"`
	WOrderId         string `json:"w_order_id,omitempty"`
	Uid              string `json:"uid,omitempty"`
	Amount           string `json:"amount,omitempty"`
	Unit             string `json:"unit,omitempty"`
	PayeeEmail       string `json:"payee_email,omitempty"`
	PayeePhoneNumber string `json:"payee_phone_number,omitempty"`
	CallbackUrl      string `json:"callback_url,omitempty"`
	CardType         string `json:"card_type,omitempty"`
	WithdrawChannel  string `json:"withdraw_channel,omitempty"`
	BankAccount      string `json:"bank_account,omitempty"`
	Ifsc             string `json:"ifsc,omitempty"`
	AccountName      string `json:"account_name,omitempty"`
	UpiId            string `json:"upi_id,omitempty"`
	Mid              string `json:"mid,omitempty"`
	Checksum         string `json:"checksum,omitempty"`
	Location         string `json:"location,omitempty"`
}
