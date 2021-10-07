package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	file := excelize.NewFile()
	sheetName := "Sheet1"

	column := 'A'
	index := 0

	file.MergeCell(sheetName, fmt.Sprintf("%c%d", column, index+1), fmt.Sprintf("%c%d", column, index+2))
	file.SetCellValue(sheetName, fmt.Sprintf("%c%d", column, index+1), "Наименование")
	column += 1

	file.MergeCell(sheetName, fmt.Sprintf("%c%d", column, index+1), fmt.Sprintf("%c%d", column, index+2))
	file.SetCellValue(sheetName, fmt.Sprintf("%c%d", column, index+1), "Артикул")
	column += 1

	file.MergeCell(sheetName, fmt.Sprintf("%c%d", column, index+1), fmt.Sprintf("%c%d", column+3, index+1))
	file.SetCellValue(sheetName, fmt.Sprintf("%c%d", column, index+1), "Import")

	if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
