package helpers

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
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := r.Scan(columnPointers...); err != nil {
			LogError("Error occurred while scanning results", err)
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
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
