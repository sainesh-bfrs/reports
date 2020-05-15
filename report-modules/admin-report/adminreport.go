package adminreport

/*
 * File: adminreport.go
 * File Created: Thursday, 14th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"fmt"
	"reports/report-modules/worker"
	"reports/services/database"
	"sync"
)

var options interface{}

// Handle ...
func Handle() []map[string]interface{} {

	fmt.Println("printing : ", options)

	type mapper struct {
		CompanyID uint8
	}

	var m []mapper
	database.DB.Table("products").Select("DISTINCT(company_id)").Scan(&m)
	var cids []uint8
	for _, company := range m {
		cids = append(cids, company.CompanyID)
	}

	a := make(chan []map[string]interface{})
	var wg sync.WaitGroup
	var data [][]map[string]interface{}
	for i, cid := range cids {
		wg.Add(1)
		go worker.RawQueryWorker(i, fmt.Sprintf("SELECT * FROM companies WHERE id = %v", cid), a, &wg)
		data = append(data, <-a)
	}
	wg.Wait()
	var total []map[string]interface{}
	for _, d := range data {
		for _, e := range d {
			total = append(total, e)
		}
	}
	return total
}

// SetOptions ...
func SetOptions(o interface{}) {
	options = o
}
