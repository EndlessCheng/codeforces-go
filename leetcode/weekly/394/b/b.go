package main

import "math/bits"

// https://space.bilibili.com/206214
func numberOfSpecialChars(word string) int {
	var lower, upper, invalid uint
	for _, c := range word {
		bit := uint(1) << (c & 31)
		if c&32 > 0 { // 小写字母
			lower |= bit
			if upper&bit > 0 {
				invalid |= bit
			}
		} else { // 大写字母
			upper |= bit
		}
	}
	// 从交集 lower & upper 中去掉不合法的字母 invalid
	return bits.OnesCount(lower & upper &^ invalid)
}

func numberOfSpecialChars2(word string) (ans int) {
	state := [27]int{}
	for _, c := range word {
		x := c & 31 // 转成数字 1~26
		if c&32 > 0 { // 小写字母
			if state[x] == 0 {
				state[x] = 1
			} else if state[x] == 2 { // 大写的后面不能有小写
				state[x] = -1
				ans--
			}
		} else { // 大写字母
			if state[x] == 0 { // 还没遇到小写，就先遇到大写了
				state[x] = -1
			} else if state[x] == 1 {
				state[x] = 2
				ans++
			}
		}
	}
	return
}
