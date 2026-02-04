package permutationdifferencebetweentwostrings

func findPermutationDifference(s string, t string) int {
    sIndexes := [26]int{}
    for i, chr := range s {
        sIndexes[chr-'a'] = i
    }

    diff := 0
    for i, chr := range t {
        chCode := chr-'a'
        diff += max(sIndexes[chCode]-i, i-sIndexes[chCode])
    }

    return diff
}