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
	switch report {
	case "admin-report":
		//adminReport(options)
		adminreport.Options = options
		prepareReport("admin-report", adminreport.Handle)
	}
}

func prepareReport(name string, s func() []map[string]interface{}) {
	result := s()
	helpers.WriteCSV(helpers.PrepareCSVData(result), "storage/data.csv")
}
