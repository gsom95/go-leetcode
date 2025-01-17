package neighboringbitwisexor

import "testing"

func Test_doesValidArrayExist(t *testing.T) {
	type args struct {
		derived []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				derived: []int{1, 1, 0},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				derived: []int{1, 1},
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				derived: []int{1, 0},
			},
			want: false,
		},
		{
			name: "My test 1",
			args: args{
				derived: []int{0},
			},
			want: true,
		},
		{
			name: "My test 2",
			args: args{
				derived: []int{1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesValidArrayExist(tt.args.derived); got != tt.want {
				t.Errorf("doesValidArrayExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
