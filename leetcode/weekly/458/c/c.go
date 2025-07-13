package main

// https://space.bilibili.com/206214
func processStr(s string, k int64) byte {
	n := len(s)
	sz := int64(0)
	for i, c := range s {
		if c == '*' {
			sz = max(sz-1, 0)
		} else if c == '#' {
			sz *= 2
		} else if c != '%' { // c 是字母
			sz++
		}
	}
	if k >= sz { // 下标越界
		return '.'
	}

	// 倒推
	for i := n - 1; ; i-- {
		c := s[i]
		if c == '*' {
			sz++ // 倒推过程中 sz 不会增大
		}
		if c == '#' {
			sz /= 2
			if k >= sz { // k 在复制后的右半边
				k -= sz
			}
		} else if c == '%' {
			k = sz - 1 - k
		} else {
			sz--
			if k == sz { // 找到答案
				return c
			}
		}
	}
}
