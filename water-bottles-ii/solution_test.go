package waterbottles

import "testing"

func Test_numWaterBottles(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				numBottles:  13,
				numExchange: 6,
			},
			want: 15,
		},
		{
			name: "example 2",
			args: args{
				numBottles:  10,
				numExchange: 3,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWaterBottles(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("numWaterBottles() = %v, want %v", got, tt.want)
			}
		})
	}
}
