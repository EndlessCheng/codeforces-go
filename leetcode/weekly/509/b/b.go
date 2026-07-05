package main

// https://space.bilibili.com/206214
func canMakeSubsequence1(s, t string) bool {
	//n, m := len(s), len(t)
	//// s[suf[i]:] 是 t[i:] 的子序列
	//suf := make([]int, m+1)
	//suf[m] = n
	//j := n
	//for i := m - 1; i >= 0; i-- {
	//	if t[i] == s[j-1] {
	//		j--
	//		if j == 0 { // s 已是 t 的子序列
	//			return true
	//		}
	//	}
	//	suf[i] = j
	//}
	//
	//pre := -1
	//for i, ch := range t {
	//	// 替换 ch
	//	if byte(ch) == s[pre+1] {
	//		pre++
	//	}
	//	// 现在，s[0,pre] 是 t[0,i] 的子序列
	//	if pre+2 >= suf[i] {
	//		return true
	//	}
	//}

	return false
}

func canMakeSubsequence(s, t string) bool {
	n := len(s)
	j0 := 0 // 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
	j1 := 0 // 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列
	for _, ch := range t {
		// j1 普通匹配
		if s[j1] == byte(ch) {
			j1++
		}

		// 也可以修改 s[j0] 为 ch，强行匹配
		j1 = max(j1, j0+1)

		// j0 普通匹配
		if s[j0] == byte(ch) {
			j0++
		}

		if j0 == n || j1 == n {
			// s 是 t 的子序列
			return true
		}
	}
	return false
}
