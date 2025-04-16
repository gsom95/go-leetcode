package countthenumberofgoodsubarrays

import "testing"

func Test_countGood(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
				k:    10,
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{3, 1, 4, 3, 2, 2, 4},
				k:    2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGood(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
