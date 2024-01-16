package excel_sheet_column_number

import (
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	tests := []struct {
		name        string
		columnTitle string
		want        int
	}{
		{
			name:        "Example 1",
			columnTitle: "A",
			want:        1,
		},
		{
			name:        "Example 2",
			columnTitle: "AB",
			want:        28,
		},
		{
			name:        "Example 3",
			columnTitle: "ZY",
			want:        701,
		},
		{
			name:        "Example 4",
			columnTitle: "FXSHRXW",
			want:        2147483647,
		},
		{
			name:        "Example 5",
			columnTitle: "AZ",
			want:        52,
		},
		{
			name:        "Example 6",
			columnTitle: "AAA",
			want:        703,
		},
		{
			name:        "Example 7",
			columnTitle: "AAAA",
			want:        18279,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
