package main

// github.com/EndlessCheng/codeforces-go
func addSpaces(s string, spaces []int) string {
	spaces = append(spaces, len(s)) // 这样可以在循环中处理最后一段
	ans := make([]byte, 0, len(s)+len(spaces))
	ans = append(ans, s[:spaces[0]]...)
	for i := 1; i < len(spaces); i++ {
		ans = append(ans, ' ')
		ans = append(ans, s[spaces[i-1]:spaces[i]]...)
	}
	return string(ans)
}

func addSpaces1(s string, spaces []int) string {
	ans := make([]byte, 0, len(s)+len(spaces))
	j := 0
	for i, c := range s {
		if j < len(spaces) && spaces[j] == i {
			ans = append(ans, ' ')
			j++
		}
		ans = append(ans, byte(c))
	}
	return string(ans)
}
