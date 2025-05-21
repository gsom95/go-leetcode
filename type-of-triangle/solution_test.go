package typeoftriangle

import "testing"

func Test_triangleType(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{nums: []int{3, 3, 3}},
			want: "equilateral",
		},
		{
			name: "Example 2",
			args: args{nums: []int{3, 4, 5}},
			want: "scalene",
		},
		{
			name: "my test 1",
			args: args{nums: []int{3, 3, 4}},
			want: "isosceles",
		},
		{
			name: "my test 2",
			args: args{nums: []int{3, 3, 7}},
			want: "none",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleType(tt.args.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
