package main

// github.com/EndlessCheng/codeforces-go
func makeFancyString(s string) string {
	ans := []byte{}
	cnt := 0
	for i, ch := range s {
		cnt++
		if cnt < 3 {
			ans = append(ans, byte(ch))
		}
		if i < len(s)-1 && byte(ch) != s[i+1] {
			cnt = 0 // 当前字母和下个字母不同，重置计数器
		}
	}
	return string(ans)
}
