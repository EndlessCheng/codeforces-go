package main

import "fmt"

// https://space.bilibili.com/206214
// 上界思维：设分割个数为 i，它能「容纳」多长的 message？
// 核心思路：枚举分割个数 i，不断增大容量 cap，直到 cap >= len(message) 为止，就可以分割了
func splitMessage(message string, limit int) []string {
	for i, cap, tailLen := 1, 0, 0; ; i++ {
		if i < 10 {
			tailLen = 5 // 结尾的长度
		} else if i < 100 {
			if i == 10 { cap -= 9 } // 前面的结尾的长度都 +1，那么容量就要减小
			tailLen = 7
		} else if i < 1000 {
			if i == 100 { cap -= 99 }
			tailLen = 9
		} else {
			if i == 1000 { cap -= 999 }
			tailLen = 11
		}
		if tailLen >= limit { return nil } // cap 无法增大，寄
		cap += limit - tailLen
		if cap < len(message) { continue } // 容量没有达到，继续枚举

		ans := make([]string, i)
		for j := range ans {
			tail := fmt.Sprintf("<%d/%d>", j+1, i)
			if j == i-1 {
				ans[j] = message + tail
			} else {
				m := limit - len(tail)
				ans[j] = message[:m] + tail
				message = message[m:]
			}
		}
		return ans
	}
}
