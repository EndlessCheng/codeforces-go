package main

import "strconv"

// https://space.bilibili.com/206214
var pow5 []string

func init() {
	// 预处理 2**15 以内的 5 的幂
	for p5 := 1; p5 < 1<<15; p5 *= 5 {
		pow5 = append(pow5, strconv.FormatUint(uint64(p5), 2))
	}
}

func minimumBeautifulSubstrings(s string) int {
	n := len(s)
	f := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		f[i] = n + 1
		if s[i] == '0' {
			continue
		}
		for _, t := range pow5 {
			if i+len(t) > n {
				break
			}
			if s[i:i+len(t)] == t {
				f[i] = min(f[i], f[i+len(t)]+1)
			}
		}
	}
	if f[0] > n {
		return -1
	}
	return f[0]
}

func min(a, b int) int { if b < a { return b }; return a }
