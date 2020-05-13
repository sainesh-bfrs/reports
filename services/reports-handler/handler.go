package reportshandler

import (
	"encoding/json"
	"fmt"
)

// Handle ...
func Handle(b []byte) {
	var r map[string]interface{}
	json.Unmarshal(b, &r)
	report := r["report"]
	options := r["option"]
	// fmt.Println(r, options)
	switch report {
	case "admin-report":
		adminReport(options)
	}
	// fmt.Println(report, options)
}

func adminReport(o interface{}) {
	fmt.Println(o)
}
