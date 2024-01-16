package palindromelinkedlist

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{head: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 1}}}}},
			want: true,
		},
		{
			name: "test2",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			want: false,
		},
		{
			name: "test3",
			args: args{head: nil},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
