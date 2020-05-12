package helpers

/*
 * File: csv.go
 * File Created: Monday, 11th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"encoding/csv"
	"fmt"
	"os"
)

// WriteCSV ...
func WriteCSV(data [][]string, filepath string) {
	file, err := os.Create(filepath)
	LogError("Unable to creare file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		LogError("Cannot write to file", err)
	}
}

// PrepareCSVData ..
func PrepareCSVData(data []map[string]interface{}) [][]string {
	var columns []string
	var rows []string
	var formatted [][]string
	for col := range data[0] {
		columns = append(columns, col)
	}
	formatted = append(formatted, columns)
	for _, row := range data {
		rows = make([]string, 0)
		for _, c := range columns {
			rows = append(rows, fmt.Sprintf("%v", row[c]))
		}
		formatted = append(formatted, rows)
	}
	return formatted
}
