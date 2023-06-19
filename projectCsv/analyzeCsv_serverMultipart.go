package projectCsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func AnalyzeCsvV2() {
	fmt.Println(os.Getwd())
	// 打开原始 CSV 文件
	file, err := os.Open("./projectCsv/server_multipart.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建一个 CSV Reader
	reader := csv.NewReader(file)

	// 读取 CSV 文件的所有行
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	res := make([]map[string]string, 0)
	// 对每一行的数据进行操作
	for _, row := range rows {
		// TODO: 在这里进行你的操作，比如修改某些列的数值
		tempRes := make(map[string]string)

		tempRes["cn"] = row[0]
		tempRes["en"] = row[1]
		tempRes["id"] = row[2]
		tempRes["br"] = row[3]
		tempRes["th"] = row[4]
		res = append(res, tempRes)
	}

	for _, v := range res {
		fmt.Println("{")
		for kk, vv := range v {
			fmt.Printf("\"%v\":\"%v\", \n", kk, vv)
		}
		fmt.Println("},")
	}
}
