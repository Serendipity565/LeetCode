package LeetCodeHot100

func minWindow(s string, t string) string {
	lens := len(s)
	lent := len(t)
	if lens < lent {
		return ""
	}
	m := make(map[byte]int)
	o := make(map[byte]int)
	for i := 0; i < lent; i++ {
		m[t[i]]++
	}

	check := func() bool {
		for k, v := range m {
			if o[k] != v {
				return false
			}
		}
		return true
	}

	ans := ""
	mymaxlen := lens
	left := 0
	for right := 0; right < lens; right++ {
		if _, ok := m[s[right]]; ok {
			o[s[right]]++
		}
		for check() && left <= right {
			if right-left+1 <= mymaxlen {
				mymaxlen = right - left + 1
				ans = s[left : right+1]
			}
			if _, ok := m[s[left]]; ok {
				o[s[left]]--
			}
			left++
		}
	}
	return ans
}
