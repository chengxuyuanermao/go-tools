package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/chengxuyuanermao/goTools/conv"
	"reflect"
)

func Test2() {
	b := &bytes.Buffer{}
	data := []map[string]interface{}{
		{"id": 1, "name": "fff"},
		{"id": 2, "name": "xxx"},
	}
	header := []string{"id", "name"}
	res := writeToCsv(b, data, header, false)
	fmt.Println(res)
	fmt.Println(string(b.Bytes()))
}

func writeToCsv(b *bytes.Buffer, dataList []map[string]interface{}, headers []string, addHeader bool) bool {
	// 写数据到csv文件
	w := csv.NewWriter(b)

	records := make([][]string, 0)
	if addHeader {
		records = append(records, headers)
	}
	for _, data := range dataList {
		row := make([]string, 0)
		for _, k := range headers {
			v := data[k]
			if v != nil {
				rt := reflect.TypeOf(v)
				if rt.Kind() == reflect.String {
					row = append(row, conv.ToString(v)+"\t") //防止出现数字字符变成科学计数
				} else {
					row = append(row, conv.ToString(v))
				}
			} else {
				row = append(row, "")
			}
		}
		records = append(records, row)
	}

	// WriteAll方法使用Write方法向w写入多条记录，并在最后调用Flush方法清空缓存。
	w.WriteAll(records)
	w.Flush()

	return true
}

func Test3() {
	// 创建 CSV 文件
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	// 写入 CSV 文件头部
	w.Write([]string{"ID", "Name", "Age"})
	// 写入数据
	w.Write([]string{"1", "hhhh", "12"})

	// 刷新缓冲区
	w.Flush()

	// 获取生成的 CSV 文件内容
	csvContent := b.Bytes()
	fmt.Println(string(csvContent))

}
