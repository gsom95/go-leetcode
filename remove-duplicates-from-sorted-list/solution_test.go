package remove_duplicates_from_sorted_list

import (
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
			want: []int{1, 2},
		},
		{
			name: "Example 2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 3,
								},
							},
						},
					},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "Example 3",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.args.head)
			gotSlice := make([]int, 0)
			for got != nil {
				gotSlice = append(gotSlice, got.Val)
				got = got.Next
			}
			if len(gotSlice) != len(tt.want) {
				t.Errorf("deleteDuplicates() result len = %d, want len = %d", len(gotSlice), len(tt.want))
				t.FailNow()
			}
			for i := 0; i < len(gotSlice); i++ {
				if gotSlice[i] != tt.want[i] {
					t.Errorf("deleteDuplicates() = %v, want %v", gotSlice, tt.want)
					t.FailNow()
				}
			}
		})
	}
}
