package simplifypath

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "My test 1",
			args: args{
				path: "/",
			},
			want: "/",
		},
		{
			name: "Example 1",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "Example 2",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
		{
			name: "Example 3",
			args: args{
				path: "/home/user/Documents/../Pictures",
			},
			want: "/home/user/Pictures",
		},
		{
			name: "Example 4",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "Example 5",
			args: args{
				path: "/.../a/../b/c/../d/./",
			},
			want: "/.../b/d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %q, want %q", got, tt.want)
			}
		})
	}
}
