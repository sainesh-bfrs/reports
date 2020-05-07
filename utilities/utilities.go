package utilities

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func mapBytesToString(m map[string]interface{}) {
	for k, v := range m {
		if b, ok := v.([]byte); ok {
			m[k] = string(b)
		}
	}
}

// MapScanExtended ...
func MapScanExtended(r *sqlx.Rows) []map[string]interface{} {
	var a []map[string]interface{}
	for r.Next() {
		results := make(map[string]interface{})
		err := r.MapScan(results)
		if err != nil {
			log.Fatalln(err)
		}
		a = append(a, results)
	}
	for _, s := range a {
		mapBytesToString(s)
	}
	return a
}
