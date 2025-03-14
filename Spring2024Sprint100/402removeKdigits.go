package Spring2024Sprint100

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := range num {
		if len(stack) == 0 || stack[len(stack)-1] <= num[i] {
			stack = append(stack, num[i])
		} else {
			for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
				stack = stack[:len(stack)-1]
				k--
			}
			stack = append(stack, num[i])
		}
	}
	cnt := 0
	// 找到第一个非0的位置
	for cnt < len(stack) {
		if stack[cnt] == '0' {
			cnt++
		} else {
			break
		}
	}
	if cnt >= len(stack) {
		return "0"
	} else {
		stack = stack[cnt:]
	}
	if len(stack) == 0 || k >= len(stack) {
		return "0"
	}
	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	return string(stack)
}
