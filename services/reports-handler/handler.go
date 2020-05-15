package reportshandler

/*
 * File: handler.go
 * File Created: Wednesday, 13th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"encoding/json"
	"reports/helpers"
	adminreport "reports/report-modules/admin-report"
)

// Handle ...
func Handle(b []byte) {
	var r map[string]interface{}
	json.Unmarshal(b, &r)
	report := r["report"]
	options := r["option"]
	var handle func() []map[string]interface{}
	switch report {
	case "admin-report":
		adminreport.SetOptions(options)
		handle = adminreport.Handle
		break
	}
	prepareReport(report.(string), handle)
}

func prepareReport(name string, s func() []map[string]interface{}) {
	result := s()
	helpers.WriteCSV(helpers.PrepareCSVData(result), "storage/data.csv")
}
