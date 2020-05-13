package helpers

/*
 * File: sqlmapper.go
 * File Created: Monday, 11th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"database/sql"
)

func mapBytesToString(m map[string]interface{}) {
	for k, v := range m {
		if b, ok := v.([]byte); ok {
			m[k] = string(b)
		}
	}
}

// MapScan ...
func MapScan(r *sql.Rows) []map[string]interface{} {

	cols, _ := r.Columns()
	var mapper []map[string]interface{}

	mapper = make([]map[string]interface{}, 0)

	for r.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := r.Scan(columnPointers...); err != nil {
			LogError("Error occurred while scanning results", err)
		}
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		mapBytesToString(m)
		mapper = append(mapper, m)
	}
	return mapper
}
