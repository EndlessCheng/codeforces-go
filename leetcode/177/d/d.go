package main

import "strings"

func largestMultipleOfThree(digits []int) (ans string) {
	cnt := [10]int{}
	sum := 0
	for _, d := range digits {
		cnt[d]++
		sum += d
	}
	del := func(v int) bool {
		for ; v < 10; v += 3 {
			if cnt[v] > 0 {
				cnt[v]--
				return true
			}
		}
		return false
	}
	if v := sum % 3; v > 0 && !del(v) {
		del(3 - v)
		del(3 - v)
	}
	for i := 9; i >= 0; i-- {
		ans += strings.Repeat(string(byte('0'+i)), cnt[i])
	}
	if ans != "" && ans[0] == '0' {
		return "0"
	}
	return
}
