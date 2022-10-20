package goXlsx

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func Use() {
	//create()
	read()
}

// 生成 xlsx
func create() {
	xlsx := excelize.NewFile()
	// Create a new sheet.
	index := xlsx.NewSheet("Sheet2")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func read() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Sheet1 has all rows len：%v， value:%v \n", len(rows), rows)
	for k1, row := range rows {
		for k2, colCell := range row {
			fmt.Printf("k1:%v,k2:%v, value:%v\n", k1, k2, colCell)
		}
	}
}
