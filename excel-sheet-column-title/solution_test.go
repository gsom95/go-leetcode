package excel_sheet_column_title

import (
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	tests := []struct {
		name         string
		columnNumber int
		want         string
	}{
		{
			name:         "Example 1",
			columnNumber: 1,
			want:         "A",
		},
		{
			name:         "Example 2",
			columnNumber: 28,
			want:         "AB",
		},
		{
			name:         "Example 3",
			columnNumber: 701,
			want:         "ZY",
		},
		{
			name:         "Example 4",
			columnNumber: 2147483647,
			want:         "FXSHRXW",
		},
		{
			name:         "Example 5",
			columnNumber: 52,
			want:         "AZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
