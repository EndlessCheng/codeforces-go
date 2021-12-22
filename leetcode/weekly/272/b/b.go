package main

// Go 模拟

// github.com/EndlessCheng/codeforces-go
func addSpaces(s string, spaces []int) string {
	spaces = append(spaces, len(s)) // 小技巧：把 s 长度加到数组末尾，这样可以在循环内处理最后一段
	ans := []byte(s[:spaces[0]])
	for i := 1; i < len(spaces); i++ {
		ans = append(ans, ' ')
		ans = append(ans, s[spaces[i-1]:spaces[i]]...)
	}
	return string(ans)
}

// github.com/EndlessCheng/codeforces-go
func addSpaces2(s string, spaces []int) string {
	ans := make([]byte, 0, len(s)+len(spaces))
	ans = append(ans, s[:spaces[0]]...)
	for i := 1; i < len(spaces); i++ {
		ans = append(ans, ' ')
		ans = append(ans, s[spaces[i-1]:spaces[i]]...)
	}
	ans = append(ans, ' ')
	ans = append(ans, s[spaces[len(spaces)-1]:]...)
	return string(ans)
}
