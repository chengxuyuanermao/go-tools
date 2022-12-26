package docx

/**
插件：github.com/nguyenthenguyen/docx
上传地址：https://cuttlefish.baidu.com/shopmis?_wkts_=1671617291075#/taskCenter/majorTask
过滤：ignore_word := ["/","党","新建文档","《","》",":","*","<",">","|","?","."]

请用以下题目《用例的作用》，写一篇320字以上的作文，并且分成至少3个段落
*/

import (
	"encoding/json"
	"fmt"
	"github.com/nguyenthenguyen/docx"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

// Use
/**
titles := []string{"噪音扰民打110不处理怎么办", "新闻播报的技巧", "括号的用法", "倒时方程的应用条件", "循环结构的概念", "电器带不动国网"}
cid从左到右，从0开始数的tab分类；i=0代表第一页
cid= 0学前教育；6互联网；7行业资料；10实用模版
done文件夹命名：
	done/日期-种类/1/title.docx
*/
func Use() {
	cid := 7
	perNum := 20
	for i := 0; i < 10; i++ { // 页数偏移
		titles := getTitles(cid, i, perNum)
		//titles := []string{"噪音扰民打110不处理怎么办"}
		for k, title := range titles {
			fmt.Printf("第%v篇开始---\n", k+1+i*perNum)
			content := getContent(title)

			r, err := docx.ReadDocxFile("./docx/pattern/template.docx")
			if err != nil {
				panic(err)
			}
			dir := fmt.Sprintf("./docx/done/%v-%v/%v/", time.Now().Format("20060102"), cid, i+1)
			// 判断文件夹是否存在，不存在要创建
			if !pathExists(dir) {
				err = os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
			}

			docx1 := r.Editable()
			docx1.Replace("content", content, -1)
			docx1.WriteToFile(dir + title + ".docx")
			r.Close()
			fmt.Printf("第%v篇结束---\n", k+1+i*perNum)
			time.Sleep(2 * time.Second)
		}
	}

}

func getContent(title string) string {
	reqTitle := fmt.Sprintf("用题目《%v》，写一篇320字以上的作文，要求分成至少3个段落，且一次性返回", title)
	content, _ := Completions(reqTitle)
	//敏感词替换
	ignoreWord := []string{"/", "党", "新建文档", "《", "》", ":", "*", "<", ">", "|", "?", "."}
	for _, word := range ignoreWord {
		content = strings.Replace(content, word, "", -1)
	}
	content = strings.TrimSpace(content)
	content = strings.TrimLeft(content, "。")
	content = strings.TrimLeft(content, "，")
	return content
}

func getTitles(cid int, page int, perNum int) []string {
	// curl 'https://cuttlefish.baidu.com/user/interface/getquerypacklist?cid=1&pn=0&rn=20&word=&tab=1' \
	//   -H 'Accept: application/json, text/plain, */*' \
	//   -H 'Accept-Language: zh-CN,zh;q=0.9' \
	//   -H 'Cache-Control: no-cache, no-store' \
	//   -H 'Connection: keep-alive' \
	//   -H 'Cookie: BIDUPSID=D8495263166146BB913F67E83F513762; PSTM=1657514247; H_WISE_SIDS=107320_110085_179347_180638_188745_194519_194530_196426_197711_199577_204907_208721_209204_209568_210305_210321_212295_212873_213032_213352_214115_214130_214137_214143_214396_214793_215108_215176_215727_216047_216212_216619_216839_216883_216941_217167_217710_218445_218459_218548_218598_218855_219156_219360_219363_219447_219451_219549_219562_219593_219666_219671_219732_219733_219742_219814_219943_219946_220068_220071_220279_220301_220340_220395_220607_220663_220774_220801_221019_221107_221117_221118_221120_221183_221369_221371_221385_221434_221467_221501_221548_221624_221717_221825_221872_221894; H_WISE_SIDS_BFESS=107320_110085_179347_180638_188745_194519_194530_196426_197711_199577_204907_208721_209204_209568_210305_210321_212295_212873_213032_213352_214115_214130_214137_214143_214396_214793_215108_215176_215727_216047_216212_216619_216839_216883_216941_217167_217710_218445_218459_218548_218598_218855_219156_219360_219363_219447_219451_219549_219562_219593_219666_219671_219732_219733_219742_219814_219943_219946_220068_220071_220279_220301_220340_220395_220607_220663_220774_220801_221019_221107_221117_221118_221120_221183_221369_221371_221385_221434_221467_221501_221548_221624_221717_221825_221872_221894; BDUSS=YxMlpUWUQ0M1NUM3NQLUtNOVpYb3lveXo2MWJQZG5KYUU5eG5vTFAwaEFaaGxqSVFBQUFBJCQAAAAAAAAAAAEAAACmXi9xuf65~rn-MjAxNsDW1LAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEDZ8WJA2fFiel; BDUSS_BFESS=YxMlpUWUQ0M1NUM3NQLUtNOVpYb3lveXo2MWJQZG5KYUU5eG5vTFAwaEFaaGxqSVFBQUFBJCQAAAAAAAAAAAEAAACmXi9xuf65~rn-MjAxNsDW1LAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEDZ8WJA2fFiel; BAIDUID=CDE23D8F39CB6E329C780AE6E5C4A01A:SL=0:NR=10:FG=1; MCITY=-257%3A; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BAIDUID_BFESS=CDE23D8F39CB6E329C780AE6E5C4A01A:SL=0:NR=10:FG=1; delPer=0; PSINO=6; BA_HECTOR=a080202h2l20ag8401ah807m1hq5li51j; ZFY=mRl0qkowq6GN8x4lf5xpn:BokGBHdm4ao1jSbTZTZyWU:C; session_id=1671616524163; session_name=cuttlefish.baidu.com; H_PS_PSSID=36555_37972_37647_37962_37909_37883_37799_37927_37900_26350_22160_37881; ZD_ENTRY=baidu; ab_sr=1.0.1_M2ViZWMzNzZiMzNlNTdhZjRiYWRiNzYwMGY3MjQ0ZDBmOTU2ZDY2MDVhM2E5OGY3NzhjYTFiNGQ2Y2ZkZDRlYjhlNDFjOGE4OTI3MmZlODVkZGM5MDgxNWVjMDE4YWE2NzRkNjkzYTQ1ZDJjM2U4MTA0ZDFjNmFkYWVkMzc2YjVhYzllMzkyODY3OWMwNWYyZTc1NWNlNDUwZTc0ZDk0ZTg4Y2RiYzRlOWNmOTViODc2Y2ZlY2QxZWNhOTdhZDY5' \
	//   -H 'Pragma: no-cache' \
	//   -H 'Referer: https://cuttlefish.baidu.com/shopmis?_wkts_=1671690689072' \
	//   -H 'Sec-Fetch-Dest: empty' \
	//   -H 'Sec-Fetch-Mode: cors' \
	//   -H 'Sec-Fetch-Site: same-origin' \
	//   -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36' \
	//   -H 'sec-ch-ua: "Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'sec-ch-ua-platform: "macOS"' \
	//   --compressed
	//

	url := fmt.Sprintf("https://cuttlefish.baidu.com/user/interface/getquerypacklist?cid=%v&pn=%v&rn=%v&word=&tab=1", cid, page, perNum)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	//  cookie设置
	req.Header.Set("Cookie", "BIDUPSID=D8495263166146BB913F67E83F513762; PSTM=1657514247; H_WISE_SIDS=107320_110085_179347_180638_188745_194519_194530_196426_197711_199577_204907_208721_209204_209568_210305_210321_212295_212873_213032_213352_214115_214130_214137_214143_214396_214793_215108_215176_215727_216047_216212_216619_216839_216883_216941_217167_217710_218445_218459_218548_218598_218855_219156_219360_219363_219447_219451_219549_219562_219593_219666_219671_219732_219733_219742_219814_219943_219946_220068_220071_220279_220301_220340_220395_220607_220663_220774_220801_221019_221107_221117_221118_221120_221183_221369_221371_221385_221434_221467_221501_221548_221624_221717_221825_221872_221894; H_WISE_SIDS_BFESS=107320_110085_179347_180638_188745_194519_194530_196426_197711_199577_204907_208721_209204_209568_210305_210321_212295_212873_213032_213352_214115_214130_214137_214143_214396_214793_215108_215176_215727_216047_216212_216619_216839_216883_216941_217167_217710_218445_218459_218548_218598_218855_219156_219360_219363_219447_219451_219549_219562_219593_219666_219671_219732_219733_219742_219814_219943_219946_220068_220071_220279_220301_220340_220395_220607_220663_220774_220801_221019_221107_221117_221118_221120_221183_221369_221371_221385_221434_221467_221501_221548_221624_221717_221825_221872_221894; BDUSS=YxMlpUWUQ0M1NUM3NQLUtNOVpYb3lveXo2MWJQZG5KYUU5eG5vTFAwaEFaaGxqSVFBQUFBJCQAAAAAAAAAAAEAAACmXi9xuf65~rn-MjAxNsDW1LAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEDZ8WJA2fFiel; BDUSS_BFESS=YxMlpUWUQ0M1NUM3NQLUtNOVpYb3lveXo2MWJQZG5KYUU5eG5vTFAwaEFaaGxqSVFBQUFBJCQAAAAAAAAAAAEAAACmXi9xuf65~rn-MjAxNsDW1LAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEDZ8WJA2fFiel; BAIDUID=CDE23D8F39CB6E329C780AE6E5C4A01A:SL=0:NR=10:FG=1; MCITY=-257%3A; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BAIDUID_BFESS=CDE23D8F39CB6E329C780AE6E5C4A01A:SL=0:NR=10:FG=1; delPer=0; PSINO=6; BA_HECTOR=a080202h2l20ag8401ah807m1hq5li51j; ZFY=mRl0qkowq6GN8x4lf5xpn:BokGBHdm4ao1jSbTZTZyWU:C; session_id=1671616524163; session_name=cuttlefish.baidu.com; H_PS_PSSID=36555_37972_37647_37962_37909_37883_37799_37927_37900_26350_22160_37881; ZD_ENTRY=baidu; ab_sr=1.0.1_M2ViZWMzNzZiMzNlNTdhZjRiYWRiNzYwMGY3MjQ0ZDBmOTU2ZDY2MDVhM2E5OGY3NzhjYTFiNGQ2Y2ZkZDRlYjhlNDFjOGE4OTI3MmZlODVkZGM5MDgxNWVjMDE4YWE2NzRkNjkzYTQ1ZDJjM2U4MTA0ZDFjNmFkYWVkMzc2YjVhYzllMzkyODY3OWMwNWYyZTc1NWNlNDUwZTc0ZDk0ZTg4Y2RiYzRlOWNmOTViODc2Y2ZlY2QxZWNhOTdhZDY5")

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache, no-store")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://cuttlefish.baidu.com/shopmis?_wkts_=1671690689072")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Println(" http.DefaultClient.Do(req) error")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("io.ReadAll error")
	}
	var titleResp TitleResp
	err = json.Unmarshal(body, &titleResp)
	if err != nil {
		log.Fatalf("json.Unmarshal(body, &titleResp) is error: %v", err)
	}
	var res []string
	titleInfo := titleResp.Data.QueryList
	for _, v := range titleInfo {
		if v.Status == 1 && utf8.RuneCountInString(v.QueryName) >= 5 {
			res = append(res, strings.TrimSpace(v.QueryName))
		}
	}
	return res
}

type TitleResp struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}
type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type QueryList struct {
	QueryID        string `json:"queryId"`
	QueryName      string `json:"queryName"`
	Status         int    `json:"status"`
	EstimatedPrice string `json:"estimatedPrice"`
}
type Data struct {
	Errstr            string      `json:"errstr"`
	Total             int         `json:"total"`
	QueryList         []QueryList `json:"queryList"`
	TaskUserNum       int         `json:"taskUserNum"`
	UserFinishTaskNum int         `json:"userFinishTaskNum"`
	UploadDayNum      int         `json:"uploadDayNum"`
	NextLevelNeedDay  int         `json:"nextLevelNeedDay"`
	NextLevelNeedTask int         `json:"nextLevelNeedTask"`
	NextLevelAward    string      `json:"nextLevelAward"`
	IsTopReward       bool        `json:"isTopReward"`
}

// 判断所给路径文件/文件夹是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Demo() {
	// Read from docx file
	r, err := docx.ReadDocxFile("./docx/TestDocument.docx")
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)

	// Or read from a filesystem object:
	// r, err := docx.ReadDocxFromFS(file string, fs fs.FS)

	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	// Replace like https://golang.org/pkg/strings/#Replace
	docx1.Replace("old_1_1", "new_1_1", -1)
	docx1.Replace("old_1_2", "new_1_2", -1)
	docx1.ReplaceLink("http://example.com/", "https://github.com/nguyenthenguyen/docx", 1)
	docx1.ReplaceHeader("out with the old", "in with the new")
	docx1.ReplaceFooter("Change This Footer", "new footer")
	docx1.WriteToFile("./new_result_1.docx")

	docx2 := r.Editable()
	docx2.Replace("old_2_1", "new_2_1", -1)
	docx2.Replace("old_2_2", "new_2_2", -1)
	docx2.WriteToFile("./new_result_2.docx")

	r.Close()
}
