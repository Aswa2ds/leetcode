package codetop

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	maxLen := 0
	chMap := map[byte]bool{}

	for right < len(s) {
		if _, ok := chMap[s[right]]; ok {
			for {
				delete(chMap, s[left])
				left++
				if s[left-1] == s[right] {
					break
				}
			}
		}
		chMap[s[right]] = true
		tmpLen := right - left + 1
		if tmpLen > maxLen {
			maxLen = tmpLen
		}
		right++
	}
	return maxLen
}
