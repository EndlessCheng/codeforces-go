package main

import "strings"

// https://space.bilibili.com/206214
func longestBalanced(s string) (ans int) {
	total0 := strings.Count(s, "0")
	total1 := len(s) - total0

	pos := map[int][]int{0: {-1}} // 见 525 题
	sum := 0 // 前缀和
	for i, ch := range s {
		sum += int(ch-'0')*2 - 1

		if p := pos[sum]; len(p) < 2 {
			pos[sum] = append(p, i)
		}

		// 不交换
		ans = max(ans, i-pos[sum][0])

		// 交换子串内的一个 1 和子串外的一个 0
		if p, ok := pos[sum-2]; ok {
			if (i-p[0]-2)/2 < total0 {
				ans = max(ans, i-p[0])
			} else if len(p) > 1 {
				ans = max(ans, i-p[1])
			}
		}

		// 交换子串内的一个 0 和子串外的一个 1
		if p, ok := pos[sum+2]; ok {
			if (i-p[0]-2)/2 < total1 {
				ans = max(ans, i-p[0])
			} else if len(p) > 1 {
				ans = max(ans, i-p[1])
			}
		}
	}
	return
}
