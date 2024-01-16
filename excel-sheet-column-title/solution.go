package excel_sheet_column_title

func convertToTitle(columnNumber int) string {
	result := make([]byte, 0, 8)
	for columnNumber > 0 {
		columnNumber--
		result = append(result, byte(columnNumber%26)+'A')
		columnNumber /= 26
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
