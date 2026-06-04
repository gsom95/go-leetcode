package exclusivetimeoffunctions

import (
	"reflect"
	"testing"
)

func Test_exclusiveTime(t *testing.T) {
	type args struct {
		n    int
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				n: 2,
				logs: []string{
					"0:start:0",
					"1:start:2",
					"1:end:5",
					"0:end:6",
				},
			},
			want: []int{3, 4},
		},
		{
			name: "test case 2",
			args: args{
				n: 1,
				logs: []string{
					"0:start:0",
					"0:end:0",
				},
			},
			want: []int{1},
		},
		{
			name: "test case 3",
			args: args{
				n: 1,
				logs: []string{
					"0:start:0",
					"0:start:2",
					"0:end:5",
					"0:start:6",
					"0:end:6",
					"0:end:7",
				},
			},
			want: []int{8},
		},
		{
			name: "test case 4",
			args: args{
				n: 2,
				logs: []string{
					"0:start:0",
					"0:start:2",
					"0:end:5",
					"1:start:6",
					"1:end:6",
					"0:end:7",
				},
			},
			want: []int{7, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exclusiveTime(tt.args.n, tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exclusiveTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
