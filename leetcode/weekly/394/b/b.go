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
		x := c & 31
		if c&32 > 0 {
			if state[x] == 0 {
				state[x] = 1
			} else if state[x] == 2 {
				state[x] = -1
				ans--
			}
		} else {
			if state[x] == 0 {
				state[x] = -1
			} else if state[x] == 1 {
				state[x] = 2
				ans++
			}
		}
	}
	return
}
