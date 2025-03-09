package LeetCodeHot100

import "sort"

type str struct {
	s     string
	sortS string
}

func groupAnagrams(strs []string) [][]string {
	strStruct := make([]str, len(strs))
	for i := 0; i < len(strs); i++ {
		strSlice := []rune(strs[i])

		sort.Slice(strSlice, func(i, j int) bool {
			return strSlice[i] < strSlice[j]
		})

		sortedStr := string(strSlice)
		strStruct[i] = str{strs[i], sortedStr}
	}

	sort.Slice(strStruct, func(i, j int) bool {
		return strStruct[i].sortS < strStruct[j].sortS
	})

	var m [][]string
	var temp []string
	temp = append(temp, strStruct[0].s)
	for i := 1; i < len(strStruct); i++ {
		if strStruct[i].sortS == strStruct[i-1].sortS {
			temp = append(temp, strStruct[i].s)
		} else {
			m = append(m, temp)
			temp = []string{}
			temp = append(temp, strStruct[i].s)
		}
	}
	m = append(m, temp)
	return m
}
