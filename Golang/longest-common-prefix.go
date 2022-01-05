package main

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := ""
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return prefix
			}
		}
		prefix += string(strs[0][i])
	}

}

func main() {
	print(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
