package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"time"
)

func Use() {
	// 不存在则创建;存在则清空;读写模式;
	file, err := os.Create("person_list.csv")
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	// 延迟关闭
	defer file.Close()

	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(file)
	// 写入数据
	w.Write([]string{"老师编号", "老师姓名", "老师特长", "日期"}) //Write 进行换行
	w.Write([]string{"t1", "老师1", "纯阳无极功", ` ` + time.Now().Format("2006-01-02 15:04:05")})
	w.Flush() // 从内存刷进文件

	// Map写入
	m := make(map[int][]string)
	m[0] = []string{"学生编号", "学生姓名", "学生特长"}
	m[1] = []string{"s1", "学生1", "乾坤大挪移"}
	m[2] = []string{"s2", "学生2", "乾坤大挪移"}
	m[3] = []string{"s3", "学生3", "乾坤大挪移"}
	m[4] = []string{"s4", "学生4", "乾坤大挪移"}
	m[5] = []string{"s5", "学生5", "乾坤大挪移"}
	m[6] = []string{"s6", "学生6", "乾坤大挪移"}
	m[7] = []string{"s7", "学生7", "乾坤大挪移"}
	m[8] = []string{"s8", "学生8", "乾坤大挪移"}
	m[9] = []string{"s9", "学生9", "乾坤大挪移"}
	m[10] = []string{"s10", "学生10", "乾坤大挪移"}

	// 按照key排序
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		w.Write(m[key])
		// 刷新缓冲
		w.Flush()
	}

}
