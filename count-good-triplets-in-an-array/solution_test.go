package countgoodtripletsinanarray

import "testing"

func Test_goodTriplets(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example 1",
			args: args{
				nums1: []int{2, 0, 1, 3},
				nums2: []int{0, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				nums1: []int{4, 0, 1, 3, 2},
				nums2: []int{4, 1, 0, 2, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodTriplets(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("goodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
