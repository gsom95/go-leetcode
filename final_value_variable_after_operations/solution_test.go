package finalvaluevariableafteroperations

import "testing"

func Test_finalValueAfterOperations(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "example 1",
			args: []string{"--X", "X++", "X++"},
			want: 1,
		},
		{
			name: "example 2",
			args: []string{"++X", "++X", "X++"},
			want: 3,
		},
		{
			name: "example 3",
			args: []string{"X++", "++X", "--X", "X--"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalValueAfterOperations(tt.args); got != tt.want {
				t.Errorf("finalValueAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
