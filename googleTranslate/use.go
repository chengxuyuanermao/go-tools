package googleTranslate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	[[["this is testing","这是一个测试",null,null,3,null,null,[[]],[[["2a0054a65181c79d122abe05e57adef1","zh_en_2021q4.md"]]]]],null,"zh-CN",null,null,null,null,[]]
	*/
	//return string(bs), nil

	// 1 反序列化。由于 json 中同一维度既含有字符串，又含有数组，为复合型的，所以比较麻烦。
	var str []interface{}
	err = json.Unmarshal(bs, &str)
	if err != nil {
		fmt.Println("err: ", err)
	}
	arr1 := str[0].([]interface{}) // 三维数组
	arr2 := arr1[0].([]interface{})
	return arr2[0].(string), nil

	// 2 直接字符串拆解
	//ss := string(bs)
	//ss = strings.ReplaceAll(ss, "[", "")
	//ss = strings.ReplaceAll(ss, "]", "")
	//ss = strings.ReplaceAll(ss, "null,", "")
	//ss = strings.Trim(ss, `"`)
	//ps := strings.Split(ss, `","`)
	//return ps[0], nil
}

func use() {
	str, err := TranslateEn2Ch("这是一个测试")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
