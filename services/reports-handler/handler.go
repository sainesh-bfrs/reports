package reportshandler

/*
 * File: handler.go
 * File Created: Wednesday, 13th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"encoding/json"
	"fmt"
	"reports/helpers"
	"reports/services/database"
	"sync"
)

// Handle ...
func Handle(b []byte) {
	var r map[string]interface{}
	json.Unmarshal(b, &r)
	report := r["report"]
	options := r["option"]
	switch report {
	case "admin-report":
		adminReport(options)
	}
}

func adminReport(o interface{}) {
	fmt.Println(o)
	sql := map[int]string{
		0: "select id, email from users LIMIT 0, 2;",
		1: "select id, email from users LIMIT 3, 2;",
	}

	a := make(chan []map[string]interface{})

	var wg sync.WaitGroup

	var data [][]map[string]interface{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(i, sql[i], a, &wg)
		data = append(data, <-a)
	}
	wg.Wait()
	var total []map[string]interface{}
	for _, d := range data {
		for _, e := range d {
			total = append(total, e)
		}
	}
	helpers.WriteCSV(helpers.PrepareCSVData(total), "storage/data.csv")
}

func worker(id int, query string, ch chan []map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	rows, err := database.DB.Raw(query).Rows()
	helpers.LogError("Error in runnig query", err)
	res := helpers.MapScan(rows)
	ch <- res
}
