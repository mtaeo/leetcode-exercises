package main

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	closingParanthesesMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if openingParantheses, ok := closingParanthesesMap[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != openingParantheses {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

func main() {
	print(isValid("({})]"))
}
