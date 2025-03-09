package LeetCodeHot100

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	i := 0
	j := 0
	ans := 0
	m := make(map[byte]int)
	for j < n {
		if _, ok := m[s[j]]; ok {
			for i < j {
				if _, c := m[s[i]]; c {
					delete(m, s[i])
				}
				if s[i] == s[j] {
					i++
					break
				}
				i++
			}
		} else {
			m[s[j]] = 1
			j++
			ans = max(ans, j-i)
			fmt.Println(ans)
		}
	}
	return ans
}
