package solvingquestionswithbrainpower

import "testing"

func Test_mostPoints(t *testing.T) {
	type args struct {
		questions [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test 1",
			args: args{
				questions: [][]int{
					{3, 2},
					{4, 3},
					{4, 4},
					{2, 5},
				},
			},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{
				questions: [][]int{
					{1, 1},
					{2, 2},
					{3, 3},
					{4, 4},
					{5, 5},
				},
			},
			want: 7,
		},
		{
			name: "wrong answer 1",
			args: args{
				questions: [][]int{
					{21, 5},
					{92, 3},
					{74, 2},
					{39, 4},
					{58, 2},
					{5, 5},
					{49, 4},
					{65, 3},
				},
			},
			want: 157,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostPoints(tt.args.questions); got != tt.want {
				t.Errorf("mostPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
