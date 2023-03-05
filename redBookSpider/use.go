package redBookSpider

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

const GenderGirl = 1
const GenderBoy = 0

const COOKIE = "gid.sig=-S5N8ahQEadOyqq6fP_0k8QrgFD7aEKKP2Zm4mKKrAQ; smidV2=2022072222103967006373798a375c84d9953394ab18360058a76e0f7f347e0; a1=185a3e023d32afh3acxleu6f79gbydud0e9qqdxjo30000485381; webId=cb2ca96d900adfeb5ce4c74926a09aba; gid=yY20qd8JYD8SyY20qd8JqdfCfqJ0ixq0SC1djVl7KiWjkAq8Dvf7fA8884Y2qYy8i8qWKJWS; gid.sign=9FPAZS+++FqA0+XsYScgl/f5KxE=; gid.ss=gSMQ9UOnDuZwH2oRGJG6BW6e4grs67TaYpnrW+8Wmd2+jbCnVJY0vcUIRfvn2nnn; timestamp2=167350422757116a580385def6140e0bafb04167a9deb18de0f2264a597a47e; timestamp2.sig=cDUjefYZrMbu3psMekW1O65wh1IRvm4DwjFYze4Zlb4; xhsTrackerId=7822679b-7af8-448f-b9b0-f062bada7d9e; xhsTrackerId.sig=LwV4bUULW60TimL4ZLOKKLPltwB4r38MymyocXk322M; customerClientId=942963027639701; x-user-id=5cca94ce000000001101885d; xhsTracker=url=question&searchengine=google; xhsTracker.sig=G0hmoAbYkliJ6rkPr3m0vNvVCs2Hv97Ab2mTTr7U8iw; xsecappid=xhs-pc-web; extra_exp_ids=h5_230301_origin,h5_1208_exp3,h5_1130_exp1,ios_wx_launch_open_app_exp,h5_video_ui_exp3,wx_launch_open_app_duration_origin,ques_exp2; extra_exp_ids.sig=qZY2awgPEE5fOM_KvOCDa1q1qW7CxbtXehhuE3gfr_E; webBuild=1.2.1; web_session=040069b0c90a1e12ce7e644931364bee9f80c1; websectiga=f47eda31ec99545da40c2f731f0630efd2b0959e1dd10d5fedac3dce0bd1e04d; sec_poison_id=4fa94bdd-611d-4c53-85d2-0722d4f05ace"

var KEYWORDS = []string{"母胎", "桃花", "二十", "岁", "脱个单", "介绍", "父母", "妈", "crush", "男硕士", "嫁", "拍拖", "喜欢", "养鱼", "媒", "回家", "脱单", "相亲", "单身", "对象", "催婚", "95", "97", "98", "99", "96", "合适的人", "男朋友", "喜欢的人", "恋爱", "回家"}

type TargetInfo struct {
	DisplayTitle string `json:"display_title"`
	DetailUrl    string `json:"detail_url"`
	DetailUrl2   string `json:"detail_url2"`
	HomepageUrl  string `json:"homepage_url"`
	City         string `json:"city"`
}

func StartSpider() {
	for i := 0; i < 5; i++ {
		use()
		randSec := rand.Intn(20) + 50

		fmt.Printf("sleep %v to next turn \n", randSec)
		time.Sleep(time.Duration(randSec) * time.Second)
	}
}

func use() {
	res := redBookSpider()
	resInfo := &SpiderInfo{}
	err := json.Unmarshal(res, resInfo)
	if err != nil {
		panic("res not right")
	}
	fmt.Printf("spider get info: %v \n", len(resInfo.Data.Items))
	var targetInfos []*TargetInfo
	keyWords := KEYWORDS
	for k, item := range resInfo.Data.Items {
		fmt.Println()
		fmt.Printf("------------ start: %v ------------\n", k+1)
		fmt.Println("checking displayTitle-----: ", item.NoteCard.DisplayTitle)

		homepageUrl := fmt.Sprintf("https://www.xiaohongshu.com/user/profile/%v", item.NoteCard.User.UserID)
		detailUrl := fmt.Sprintf("https://www.xiaohongshu.com/explore/%v", item.ID)
		detailUrl2 := fmt.Sprintf("https://www.xiaohongshu.com/discovery/item/%v", item.ID)
		// 验证关键词是否命中
		isHitKeyWords := false
		for _, v := range keyWords {
			isContainKeyWord := strings.Contains(item.NoteCard.DisplayTitle, v)
			if !isContainKeyWord {
				continue
			}
			isHitKeyWords = true
			break
		}
		if !isHitKeyWords {
			fmt.Println("not hit keyword, filter--------")
			continue
		}

		// 城市&性别是否命中
		fmt.Println("checking city&gender ---- ")
		lock, city := isSuit(homepageUrl, detailUrl)
		if !lock {
			continue
		}
		fmt.Println("got item success-------")
		// 通过
		temp := &TargetInfo{
			DisplayTitle: item.NoteCard.DisplayTitle,
			DetailUrl:    detailUrl,
			DetailUrl2:   detailUrl2,
			HomepageUrl:  homepageUrl,
			City:         city,
		}
		targetInfos = append(targetInfos, temp)
	}
	fmt.Println("res-- : ", len(targetInfos))
	if len(targetInfos) > 0 {
		exportCsv(targetInfos)
		fmt.Println("export content success!")
		getComment(targetInfos)
	}
}

func exportCsv(targetInfos []*TargetInfo) {
	// 不存在则创建;存在则清空;读写模式;
	path := fmt.Sprintf("/Users/chenwenxing/goModProject/output/goTools/redBookSpider/data/person_list_%v.csv", time.Now().Format("0102150405"))
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	// 延迟关闭
	defer file.Close()

	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"标题", "帖子详情", "帖子详情2", "主页", "城市"}) //Write 进行换行
	for _, v := range targetInfos {
		temp := []string{v.DisplayTitle, v.DetailUrl, v.DetailUrl2, v.HomepageUrl, v.City}
		w.Write(temp)
		// 刷新缓冲
		w.Flush()
	}
}

func isSuit(url string, detailUrl string) (bool, string) {
	sec := time.Duration(rand.Intn(8) + 5)
	fmt.Println("sleep sec----", sec)
	time.Sleep(sec * time.Second)

	fmt.Println(url)
	body := getUserProfile(url, detailUrl)
	if body == nil {
		return false, ""
	}
	// 正则匹配
	// 城市
	r1 := regexp.MustCompile(`ipLocation":"(.*?)",`)
	if r1 == nil {
		fmt.Println("regexp is nil")
		return false, ""
	}
	res := r1.FindAllStringSubmatch(string(body), -1)
	if len(res) == 0 || len(res[0]) != 2 {
		fmt.Println("match error")
		return false, ""
	}
	city := res[0][1]
	fmt.Println("city---", city)
	if !strings.Contains(city, "广州") && !strings.Contains(city, "广东") {
		fmt.Println("filter-city-----")
		return false, city
	}
	fmt.Println("pass-city---", city)

	/**
		// 性别
		r2 := regexp.MustCompile(`"basicInfo":\{"gender":(\d?),`)
		if r2 == nil {
			fmt.Println("regexp is nil")
			return false, ""
		}
		res2 := r2.FindAllStringSubmatch(string(body), -1)
		if len(res2) == 0 || len(res2[0]) != 2 {
			fmt.Println(res2)
			fmt.Println("match2 error")
			return false, ""
		}
		gender, _ := strconv.Atoi(res2[0][1])
		if gender != GenderGirl {
			fmt.Println("filter-gender-----")
			return false, city
		}
		fmt.Println("pass-gender---", gender)
	**/
	return true, city
}
