package Spring2024Sprint100

func isValid(s string) bool {
	var stack []rune

	matching := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != matching[char] {
				return false
			}
		}
	}

	return len(stack) == 0
}
