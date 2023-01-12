package redBookSpider

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getUserProfile(url string, detailUrl string) []byte {
	// curl 'https://www.xiaohongshu.com/user/profile/61cd5a6c0000000010004e90' \
	//   -H 'authority: www.xiaohongshu.com' \
	//   -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
	//   -H 'accept-language: zh-CN,zh;q=0.9' \
	//   -H 'cache-control: max-age=0' \
	//   -H 'cookie: gid.sig=-S5N8ahQEadOyqq6fP_0k8QrgFD7aEKKP2Zm4mKKrAQ; smidV2=2022072222103967006373798a375c84d9953394ab18360058a76e0f7f347e0; webBuild=1.0.29; xsecappid=xhs-pc-web; a1=185a3e023d32afh3acxleu6f79gbydud0e9qqdxjo30000485381; webId=cb2ca96d900adfeb5ce4c74926a09aba; gid=yY20qd8JYD8SyY20qd8JqdfCfqJ0ixq0SC1djVl7KiWjkAq8Dvf7fA8884Y2qYy8i8qWKJWS; gid.sign=9FPAZS+++FqA0+XsYScgl/f5KxE=; gid.ss=gSMQ9UOnDuZwH2oRGJG6BW6e4grs67TaYpnrW+8Wmd2+jbCnVJY0vcUIRfvn2nnn; web_session=0400698dbed78617d80b653d8a314b8ac5d228; timestamp2=167350422757116a580385def6140e0bafb04167a9deb18de0f2264a597a47e; timestamp2.sig=cDUjefYZrMbu3psMekW1O65wh1IRvm4DwjFYze4Zlb4; xhsTrackerId=7822679b-7af8-448f-b9b0-f062bada7d9e; xhsTrackerId.sig=LwV4bUULW60TimL4ZLOKKLPltwB4r38MymyocXk322M; xhsTracker=&url=explore&searchengine=google; xhsTracker.sig=zQ3lZkVmYE96_48WhMlVnt_VgoOTOqjw6Jk1Xfb9cQE; websectiga=88777433e9b3e11299d61d9623d638c2e708325142208233522732ccb93770dc; extra_exp_ids=h5_1208_exp3,h5_1130_exp1,ios_wx_launch_open_app_exp,h5_video_ui_exp3,wx_launch_open_app_duration_origin,ques_exp2' \
	//   -H 'referer: https://www.xiaohongshu.com/explore/63a166ff0000000022038789' \
	//   -H 'sec-ch-ua: "Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'sec-ch-ua-platform: "macOS"' \
	//   -H 'sec-fetch-dest: document' \
	//   -H 'sec-fetch-mode: navigate' \
	//   -H 'sec-fetch-site: same-origin' \
	//   -H 'sec-fetch-user: ?1' \
	//   -H 'upgrade-insecure-requests: 1' \
	//   -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36' \
	//   --compressed

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
		fmt.Println("error get request")
		return nil
	}
	req.Header.Set("Authority", "www.xiaohongshu.com")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Cookie", "gid.sig=-S5N8ahQEadOyqq6fP_0k8QrgFD7aEKKP2Zm4mKKrAQ; smidV2=2022072222103967006373798a375c84d9953394ab18360058a76e0f7f347e0; webBuild=1.0.29; a1=185a3e023d32afh3acxleu6f79gbydud0e9qqdxjo30000485381; webId=cb2ca96d900adfeb5ce4c74926a09aba; gid=yY20qd8JYD8SyY20qd8JqdfCfqJ0ixq0SC1djVl7KiWjkAq8Dvf7fA8884Y2qYy8i8qWKJWS; gid.sign=9FPAZS+++FqA0+XsYScgl/f5KxE=; gid.ss=gSMQ9UOnDuZwH2oRGJG6BW6e4grs67TaYpnrW+8Wmd2+jbCnVJY0vcUIRfvn2nnn; web_session=0400698dbed78617d80b653d8a314b8ac5d228; timestamp2=167350422757116a580385def6140e0bafb04167a9deb18de0f2264a597a47e; timestamp2.sig=cDUjefYZrMbu3psMekW1O65wh1IRvm4DwjFYze4Zlb4; xhsTrackerId=7822679b-7af8-448f-b9b0-f062bada7d9e; xhsTrackerId.sig=LwV4bUULW60TimL4ZLOKKLPltwB4r38MymyocXk322M; xhsTracker=&url=explore&searchengine=google; xhsTracker.sig=zQ3lZkVmYE96_48WhMlVnt_VgoOTOqjw6Jk1Xfb9cQE; extra_exp_ids=h5_1208_exp3,h5_1130_exp1,ios_wx_launch_open_app_exp,h5_video_ui_exp3,wx_launch_open_app_duration_origin,ques_exp2; xsecappid=login; web_sec_uuid=1210c1cd-3a90-4324-b730-a7b5adf7b19f; websectiga=0f2f555ec25a4e63bf019a5a7c615b7b6050133c69c92d4274f0b500e6801365")
	req.Header.Set("Referer", detailUrl)
	req.Header.Set("Sec-Ch-Ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Println("error user profile")
		return nil
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		fmt.Println("ok--")
	} else {
		fmt.Println("error user profile not 200")
		return nil
	}
	return res
}
