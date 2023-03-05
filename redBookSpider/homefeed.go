package redBookSpider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SpiderInfo struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    Data   `json:"data"`
}
type User struct {
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	UserID   string `json:"user_id"`
}
type InteractInfo struct {
	Liked      bool   `json:"liked"`
	LikedCount string `json:"liked_count"`
}
type Cover struct {
	FileID  string `json:"file_id"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	URL     string `json:"url"`
	TraceID string `json:"trace_id"`
}
type NoteCard struct {
	Type         string       `json:"type"`
	DisplayTitle string       `json:"display_title"`
	User         User         `json:"user"`
	InteractInfo InteractInfo `json:"interact_info"`
	Cover        Cover        `json:"cover"`
}
type Items struct {
	ID        string   `json:"id"`
	ModelType string   `json:"model_type"`
	NoteCard  NoteCard `json:"note_card"`
	TrackID   string   `json:"track_id"`
}
type Data struct {
	CursorScore string  `json:"cursor_score"`
	Items       []Items `json:"items"`
}

func redBookSpider() []byte {
	// curl 'https://edith.xiaohongshu.com/api/sns/web/v1/homefeed' \
	//   -H 'authority: edith.xiaohongshu.com' \
	//   -H 'accept: application/json, text/plain, */*' \
	//   -H 'accept-language: zh-CN,zh;q=0.9' \
	//   -H 'content-type: application/json;charset=UTF-8' \
	//   -H 'cookie: gid.sig=-S5N8ahQEadOyqq6fP_0k8QrgFD7aEKKP2Zm4mKKrAQ; smidV2=2022072222103967006373798a375c84d9953394ab18360058a76e0f7f347e0; webBuild=1.0.29; xsecappid=xhs-pc-web; a1=185a3e023d32afh3acxleu6f79gbydud0e9qqdxjo30000485381; webId=cb2ca96d900adfeb5ce4c74926a09aba; gid=yY20qd8JYD8SyY20qd8JqdfCfqJ0ixq0SC1djVl7KiWjkAq8Dvf7fA8884Y2qYy8i8qWKJWS; gid.sign=9FPAZS+++FqA0+XsYScgl/f5KxE=; gid.ss=gSMQ9UOnDuZwH2oRGJG6BW6e4grs67TaYpnrW+8Wmd2+jbCnVJY0vcUIRfvn2nnn; web_session=0400698dbed78617d80b653d8a314b8ac5d228; timestamp2=167350422757116a580385def6140e0bafb04167a9deb18de0f2264a597a47e; timestamp2.sig=cDUjefYZrMbu3psMekW1O65wh1IRvm4DwjFYze4Zlb4; web_sec_uuid=df341f8b-681d-4cae-9de7-940bcaa1ad69; websectiga=0f9ee0b2c1171260c28c9568efd8c59c60e1a6d06a847141097dba3275398d82; xhsTrackerId=7822679b-7af8-448f-b9b0-f062bada7d9e; xhsTrackerId.sig=LwV4bUULW60TimL4ZLOKKLPltwB4r38MymyocXk322M; extra_exp_ids=h5_1208_exp3,h5_1130_exp1,ios_wx_launch_open_app_exp,h5_video_ui_exp3,wx_launch_open_app_duration_origin,ques_exp2; extra_exp_ids.sig=Iw6nBv0KxUOEEtlo5ci79RYBK7O-I3uwnXYWGdfO0Ec; xhsTracker=&url=explore&searchengine=google; xhsTracker.sig=zQ3lZkVmYE96_48WhMlVnt_VgoOTOqjw6Jk1Xfb9cQE' \
	//   -H 'origin: https://www.xiaohongshu.com' \
	//   -H 'referer: https://www.xiaohongshu.com/' \
	//   -H 'sec-ch-ua: "Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'sec-ch-ua-platform: "macOS"' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-site: same-site' \
	//   -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36' \
	//   -H 'x-s: Z6TCsYOkZB1+1B46O2akOYMGZjMCZ6TL0j9bOBVU0g93' \
	//   -H 'x-t: 1673504881056' \
	//   --data-raw '{"cursor_score":"","num":24,"refresh_type":1,"note_index":0,"unread_begin_note_id":"","unread_end_note_id":"","unread_note_count":0,"category":"homefeed_recommend"}' \
	//   --compressed

	type Payload struct {
		CursorScore       string `json:"cursor_score"`
		Num               int    `json:"num"`
		RefreshType       int    `json:"refresh_type"`
		NoteIndex         int    `json:"note_index"`
		UnreadBeginNoteID string `json:"unread_begin_note_id"`
		UnreadEndNoteID   string `json:"unread_end_note_id"`
		UnreadNoteCount   int    `json:"unread_note_count"`
		Category          string `json:"category"`
	}

	data := Payload{
		// fill struct
		Category:          "homefeed_recommend",
		CursorScore:       "",
		NoteIndex:         0,
		Num:               24,
		RefreshType:       1,
		UnreadBeginNoteID: "",
		UnreadEndNoteID:   "",
		UnreadNoteCount:   0,
	}
	//timeMill := strconv.Itoa(int(time.Now().UnixMilli()))
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://edith.xiaohongshu.com/api/sns/web/v1/homefeed", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "edith.xiaohongshu.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", COOKIE)
	req.Header.Set("Origin", "https://www.xiaohongshu.com")
	req.Header.Set("Referer", "https://www.xiaohongshu.com/")
	req.Header.Set("Sec-Ch-Ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("X-S", "Z6TCsYOkZB1+1B46O2akOYMGZjMCZ6TL0j9bOBVU0g93")
	req.Header.Set("X-T", "1673504881056")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	} else {
		fmt.Println(resp.StatusCode)
		panic("not 200")
	}
	return res
}
