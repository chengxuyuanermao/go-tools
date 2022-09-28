package googleTranslate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func TranslateEn2Ch(text string) (string, error) {
	/**
	client=gtx
	&sl=zh-cn   源语言
	&tl=en 		目标语言
	&dt=t
	&q=%s  查询参数
	*/
	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=zh-cn&tl=en&dt=t&q=%s", url.QueryEscape(text))
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 返回的bs是一个json格式的数据，用的是 byte切片存储
	/**
	[[["Programmer Er Mao is a good public account","程序员二毛是个不错的公众号",null,null,3,null,null,[[]],[[["2a0054a65181c79d122abe05e57adef1","zh_en_2021q4.md"]]]]],null,"zh-CN",null,null,null,null,[]]
	*/
	//return string(bs), nil

	// 1 反序列化。由于 json 中同一维度既含有字符串，又含有数组，为复合型的，所以比较麻烦。

	// 2 直接字符串拆解
	ss := string(bs)
	ss = strings.ReplaceAll(ss, "[", "")
	ss = strings.ReplaceAll(ss, "]", "")
	ss = strings.ReplaceAll(ss, "null,", "")
	ss = strings.Trim(ss, `"`)
	ps := strings.Split(ss, `","`)
	return ps[0], nil
}

func use() {
	str, err := TranslateEn2Ch("程序员二毛是个不错的公众号")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

type translateResp struct {
	param0 [][]interface{}
	param1 interface{}
	param2 interface{}
	param3 interface{}
	param4 interface{}
	param5 interface{}
	param6 interface{}
	param7 interface{}
}

type translateResps struct {
	translateResp []translateResp
}
