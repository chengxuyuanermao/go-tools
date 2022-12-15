package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type StockDetail struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	Ratio         string `json:"ratio"`
	Increase      string `json:"increase"`
	Volume        string `json:"volume"`
	TurnoverRatio string `json:"turnoverRatio"`
}

type Stock struct {
	Stock []StockDetail `json:"Stock"`
}

type StockResponse struct {
	Result Stock `json:"Result"`
}

func HttpGetRequest(url string) (rs []byte, err error) {
	// http.Get在net/http中，所以要import "net/http"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func main() {
	res := GetStock("苏州股票", "股票")
	fmt.Println(res)
}

func GetStock(stock string, isStockContain string) string {
	url := "https://finance.pae.baidu.com/selfselect/sug?wd=" + strings.Replace(stock, isStockContain, "", -1) + "&skip_login=1&finClientType=pc"
	rlt, err := HttpGetRequest(url)

	if err != nil {
		fmt.Println("net req error")
		return "查询" + isStockContain + "接口短暂离线了！请稍后重试"
	} else {
		var res StockResponse
		rlt = []byte(`{"QueryID":"0","ResultCode":"403","Result":{"Stock": [] }}`)

		jsonerr := json.Unmarshal(rlt, &res)
		if jsonerr != nil {
			// 解析JSON异常，根据自身业务逻辑进行调整修改
			fmt.Println("请求异常:%v", jsonerr)
			fmt.Println(string(rlt))
			return "该" + isStockContain + "不存在喔！请再搜搜"
		} else {
			var text string
			for k, v := range res.Result.Stock {
				if k <= 1 {
					text += fmt.Sprintf("\n"+isStockContain+"代码：%v\n"+isStockContain+"名称：%v\n最新价格：%v\n涨幅:%v\n涨跌额:%v\n成交量:%v\n换手率:%v\n",
						v.Code,
						v.Name,
						v.Price,
						v.Ratio,
						v.Increase,
						v.Volume,
						v.TurnoverRatio,
					)
				}
			}
			return text
		}
	}
}
