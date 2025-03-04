package lemonadechange

import "testing"

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Failed case 1",
			args: args{
				bills: []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5},
			},
			want: true,
		},
		{
			name: "Failed case 2",
			args: args{
				bills: []int{5,5,5,10,5,5,10,20,20,20},
			},
			want: false,
		},
		{
			name: "Example 1",
			args: args{
				bills: []int{5, 5, 5, 10, 20},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				bills: []int{5, 5, 10, 10, 20},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.args.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
