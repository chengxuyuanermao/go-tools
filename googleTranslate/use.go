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

	//return string(bs), nil
	//返回的json反序列化比较麻烦, 直接字符串拆解
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
