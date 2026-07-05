package main

// https://space.bilibili.com/206214
func minOperations(s1, t string) (ans int) {
	n := len(s1)
	if n == 1 && s1 == "1" && t == "0" {
		return -1
	}

	// 也可以用一个布尔变量表示 s1[i] 是否操作过，从而做到 O(1) 空间，见 Python3 写法二
	s := []byte(s1)
	for i := range n {
		if s[i] == t[i] {
			continue
		}
		if s[i] == '0' {
			ans++
		} else if i < n-1 && s[i+1] == '1' {
			ans++
			s[i+1] = '0'
		} else {
			ans += 2
		}
	}
	return
}
