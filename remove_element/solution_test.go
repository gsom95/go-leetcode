package removeelement

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	type want struct {
		num  int
		nums map[int]int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: want{
				num:  2,
				nums: map[int]int{2: 2},
			},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: want{
				num: 5,
				nums: map[int]int{
					0: 2,
					1: 1,
					3: 1,
					4: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != tt.want.num {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
				return
			}

			for i := range tt.args.nums[:got] {
				num := tt.args.nums[i]
				amount, found := tt.want.nums[num]
				if !found {
					t.Errorf("removeElement(): extra number(%d) found at position %d", num, i)
					return
				}
				if amount == 1 {
					delete(tt.want.nums, num)
					continue
				}
				tt.want.nums[num]--
			}
		})
	}
}
