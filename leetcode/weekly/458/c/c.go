package main

// https://space.bilibili.com/206214
func processStr1(s string, k int64) byte {
	n := len(s)
	size := make([]int64, n)
	sz := int64(0)
	for i, c := range s {
		if c == '*' {
			sz = max(sz-1, 0)
		} else if c == '#' {
			sz *= 2
		} else if c != '%' { // c 是字母
			sz++
		}
		size[i] = sz
	}

	if k >= size[n-1] { // 下标越界
		return '.'
	}

	// 迭代
	for i := n - 1; ; i-- {
		c := s[i]
		sz = size[i]
		if c == '#' {
			if k >= sz/2 { // k 在复制后的右半边
				k -= sz / 2
			}
		} else if c == '%' {
			k = sz - 1 - k // 反转前的下标为 sz-1-k 的字母就是答案
		} else if c != '*' && k == sz-1 { // 找到答案
			return c
		}
	}
}

func processStr(s string, k int64) byte {
	sz := int64(0)
	for _, c := range s {
		if c == '*' {
			sz = max(sz-1, 0)
		} else if c == '#' {
			sz *= 2
		} else if c != '%' {
			sz++
		}
	}

	if k >= sz {
		return '.'
	}

	for i := len(s) - 1; ; i-- {
		c := s[i]
		if c == '*' {
			sz++
		} else if c == '#' {
			sz /= 2
			if k >= sz {
				k -= sz
			}
		} else if c == '%' {
			k = sz - 1 - k
		} else {
			sz--
			if k == sz {
				return c
			}
		}
	}
}
