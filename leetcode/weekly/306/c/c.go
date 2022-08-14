package main

// https://space.bilibili.com/206214
func smallestNumber(pattern string) string {
	n := len(pattern)
	ans := make([]byte, n+1)
	for i, cur := 0, byte('1'); i < n; {
		if i > 0 && pattern[i] == 'I' {
			i++
		}
		for ; i < n && pattern[i] == 'I'; i++ {
			ans[i] = cur
			cur++
		}
		i0 := i
		for ; i < n && pattern[i] == 'D'; i++ {
		}
		for j := i; j >= i0; j-- {
			ans[j] = cur
			cur++
		}
	}
	return string(ans)
}
