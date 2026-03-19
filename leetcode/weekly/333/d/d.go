package main

// https://space.bilibili.com/206214
func findTheString(lcp [][]int) string {
	n := len(lcp)
	s := make([]byte, n)
	i := 0 // s[i] 没有填字母
	for c := byte('a'); c <= 'z'; c++ {
		for j := i; j < n; j++ {
			if lcp[i][j] > 0 { // s[j] == s[i]
				s[j] = c
			}
		}
		// 找下一个空位
		for i < n && s[i] > 0 {
			i++
		}
		if i == n { // 没有空位
			break
		}
	}

	if i < n { // 还有空位
		return ""
	}

	// 验证 s 是否符合 lcp 矩阵
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 计算后缀 [i,n-1] 和后缀 [j,n-1] 的实际 LCP
			actualLcp := 0
			if s[i] == s[j] {
				if i == n-1 || j == n-1 {
					actualLcp = 1
				} else {
					actualLcp = lcp[i+1][j+1] + 1
				}
			}
			if lcp[i][j] != actualLcp { // 矛盾
				return ""
			}
		}
	}
	return string(s)
}
