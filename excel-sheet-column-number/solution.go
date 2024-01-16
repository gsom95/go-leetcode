package excel_sheet_column_number

func titleToNumber(columnTitle string) int {
	var result int
	for _, c := range columnTitle {
		result = int(c-'A'+1) + result*26
	}
	return result
}
