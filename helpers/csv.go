package helpers

import (
	"encoding/csv"
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
