package redBookSpider

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const GenderGirl = 1
const GenderBoy = 0

type TargetInfo struct {
	DisplayTitle string `json:"display_title"`
	DetailUrl    string `json:"detail_url"`
	HomepageUrl  string `json:"homepage_url"`
	City         string `json:"city"`
}

func Use() {
	res := redBookSpider()
	resInfo := &SpiderInfo{}
	err := json.Unmarshal(res, resInfo)
	if err != nil {
		panic("res not right")
	}
	var targetInfos []*TargetInfo
	keyWords := []string{"脱单", "相亲", "单身", "对象", "催婚", "95后", "97", "98", "99", "96", "合适的人", "男朋友", "喜欢的人", "谈恋爱"}
	for _, item := range resInfo.Data.Items {
		homepageUrl := fmt.Sprintf("https://www.xiaohongshu.com/user/profile/%v", item.NoteCard.User.UserID)
		detailUrl := fmt.Sprintf("https://www.xiaohongshu.com/explore/%v", item.ID)
		lock, city := isSuit(homepageUrl, detailUrl)
		if !lock {
			continue
		}

		for _, v := range keyWords {
			isContainKeyWord := strings.Contains(item.NoteCard.DisplayTitle, v)
			if !isContainKeyWord {
				continue
			}
			temp := &TargetInfo{
				DisplayTitle: item.NoteCard.DisplayTitle,
				DetailUrl:    detailUrl,
				HomepageUrl:  homepageUrl,
				City:         city,
			}
			targetInfos = append(targetInfos, temp)
			break
		}
	}
	fmt.Println(targetInfos)
}

func isSuit(url string, detailUrl string) (bool, string) {
	time.Sleep(1 * time.Second)
	fmt.Println(url)
	body := getUserProfile(url, detailUrl)
	if body == nil {
		return false, ""
	}
	// 正则匹配
	// 城市
	r1 := regexp.MustCompile(`location":"(.*?)",`)
	if r1 == nil {
		fmt.Println("regexp is nil")
		return false, ""
	}
	res := r1.FindAllStringSubmatch(string(body), -1)
	if len(res) == 0 || len(res[0]) != 2 {
		fmt.Println(res)
		fmt.Println("match error")
		return false, ""
	}
	city := res[0][1]
	fmt.Println("city---", city)
	if !strings.Contains(city, "广州") && !strings.Contains(city, "广东") {
		return false, city
	}

	// 性别
	r2 := regexp.MustCompile(`"gender":(\d?),`)
	if r2 == nil {
		fmt.Println("regexp is nil")
		return false, ""
	}
	res2 := r2.FindAllStringSubmatch(string(body), -1)
	if len(res2) == 0 || len(res2[0]) != 2 {
		fmt.Println(res)
		fmt.Println("match2 error")
		return false, ""
	}
	gender, _ := strconv.Atoi(res2[0][1])
	if gender != GenderGirl {
		return false, city
	}

	return true, city
}
