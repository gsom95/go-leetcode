package simplifypath

import "strings"

func simplifyPath(path string) string {
	// split path by '/'
	paths := make([]string, 0)
	for _, p := range strings.Split(path, "/") {
		if p == "" || p == "." {
			continue
		}
		if p == ".." {
			if len(paths) > 0 {
				paths = paths[:len(paths)-1]
			}
			continue
		}
		paths = append(paths, p)
	}
	return "/" + strings.Join(paths, "/")
}
