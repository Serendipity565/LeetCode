package LeetCodeHot100

func findAnagrams(s string, p string) []int {
	lens, lenp := len(s), len(p)

	var ans []int

	if lens < lenp {
		return ans
	}

	pCount := [26]int{}
	for i := 0; i < lenp; i++ {
		pCount[p[i]-'a']++
	}

	sCount := [26]int{}
	for i := 0; i < lenp; i++ {
		sCount[s[i]-'a']++
	}

	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i := lenp; i < lens; i++ {
		sCount[s[i]-'a']++
		sCount[s[i-lenp]-'a']--
		if sCount == pCount {
			ans = append(ans, i-lenp+1)
		}
	}

	return ans
}
