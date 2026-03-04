package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	maxLen := 0
	seen := make(map[byte]struct{}, len(s))
	left, right := 0, 0

	for right < len(s) {
		if _, found := seen[s[right]]; found {
			delete(seen, s[left])
			left++
			continue
		}

		seen[s[right]] = struct{}{}
		right++
		maxLen = max(maxLen, right-left)
	}

	return maxLen
}
