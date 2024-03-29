package longestcommonprefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "example 1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "example 2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "my test 1",
			strs: []string{"", "racecar", "car"},
			want: "",
		},
		{
			name: "failed test 1",
			strs: []string{"ab", "a"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
