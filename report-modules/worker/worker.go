package worker

import (
	"fmt"
	"reports/helpers"
	"reports/services/database"
	"sync"
)

/*
 * File: worker.go
 * File Created: Thursday, 14th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

// RawQueryWorker ...
func RawQueryWorker(id int, query string, ch chan []map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Running worker %v", id)
	rows, err := database.DB.Raw(query).Rows()
	helpers.LogError("Error in runnig query", err)
	res := helpers.MapScan(rows)
	ch <- res
}
