package projectCsv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func AnalyzeCsv() {
	fmt.Println(os.Getwd())
	// 打开原始 CSV 文件
	file, err := os.Open("./projectCsv/act_config.csv")
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

	// 对每一行的数据进行操作
	for i, row := range rows {
		// TODO: 在这里进行你的操作，比如修改某些列的数值
		if i == 0 {
			continue
		}

		totalReward, _ := strconv.Atoi(strings.Trim(row[4], " "))
		totalReward = totalReward * 100
		ticketNumberReward, _ := strconv.Atoi(strings.Trim(row[5], " "))
		//row[0] = row[0]
		row[1] = strconv.Itoa(transTime(row[1]))
		row[2] = strconv.Itoa(transTime(row[2]))
		row[3] = strconv.Itoa(transTime(row[3]))
		row[4] = strconv.Itoa(totalReward)
		row[5] = strconv.Itoa(ticketNumberReward)
		// 输出修改后的行数据
		fmt.Printf("Row %d: %v\n", i+1, row)
	}

	// 创建一个 CSV Writer，用于写入修改后的数据
	outputFile, err := os.Create("./projectCsv/output.csv")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)

	// 将修改后的数据写入到输出文件中
	for _, row := range rows {
		err := writer.Write(row)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
}

func transTime(str string) int {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	t, err := time.ParseInLocation("2006.01.02", str, loc)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	zeroTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	return int(zeroTime.Unix())
}
