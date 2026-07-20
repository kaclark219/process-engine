package historian

import ("database/sql"
		_ "modernc.org/sqlite")

func LoadTimestamp(timestamp string) ([]map[string]any, error) {
	db, err := sql.Open("sqlite", "C:/sqlite/historian.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
			SELECT ProcessName,
				VariableName,
				Value,
				TimestampUTC
			FROM ProcessData
			WHERE TimestampUTC = ?
		`, timestamp)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []map[string]any

	for rows.Next() {
		var process string
		var variable string
		var value float64
		var ts string

		err := rows.Scan(&process, &variable, &value, &ts)
		if err != nil {
			return nil, err
		}

		records = append(records, map[string]any{
			"process": process,
			"variable": variable,
			"value": value,
			"timestamp": ts,
		})
	}

	return records, nil
}