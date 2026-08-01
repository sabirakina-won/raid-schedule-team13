package main

func validate(rows []string) string {
	if len(rows) == 0 {
		return "error: empty schedule"
	}

	if len(rows[0]) == 0 {
		return "error: empty schedule"
	}

	rowLength := len(rows[0])

	for _, row := range rows {
		if len(row) != rowLength {
			return "error: rows have different lengths"
		}

		for i := 0; i < len(row); i++ {
			symbol := row[i]

			if symbol != '0' && symbol != '1' && symbol != '2' {
				return "error: invalid symbol (only 0/1/2 allowed)"
			}
		}
	}

	return ""
}
