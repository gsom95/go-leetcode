package groupanagrams

import (
	"hash/fnv"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	words := make(map[uint64][]string, len(strs))

	hasher := fnv.New64a()
	for _, str := range strs {
		hasher.Reset()
		s := []byte(str)
		slices.Sort(s)
		hasher.Write(s)
		hashSum := hasher.Sum64()

		if other, found := words[hashSum]; found {
			other = append(other, str)
			words[hashSum] = other
		} else {
			words[hashSum] = []string{str}
		}
	}

	result := make([][]string, 0, len(words))
	for _, words := range words {
		result = append(result, words)
	}
	return result
}
