package main

// https://space.bilibili.com/206214
func smallestBeautifulString(S string, k int) string {
	limit := 'a' + byte(k)
	s := []byte(S)
	n := len(s)
	i := n - 1
	s[i]++
	for i < n {
		if s[i] == limit { // 超过范围
			if i == 0 {
				return ""
			}
			// 进位
			s[i] = 'a'
			i--
			s[i]++
		} else if i > 0 && s[i] == s[i-1] || i > 1 && s[i] == s[i-2] {
			s[i]++ // 如果和前面的形成回文串，就继续增加
		} else {
			i++ // 检查是否和后面的形成回文串
		}
	}
	return string(s)
}
