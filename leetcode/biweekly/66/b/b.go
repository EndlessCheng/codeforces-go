package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func minimumBuckets(hamsters string) int {
	if strings.Contains("H"+hamsters+"H", "HHH") {
		return -1
	}
	return strings.Count(hamsters, "H") - strings.Count(hamsters, "H.H")
}

func minimumBuckets2(hamsters string) (ans int) {
	lastFood := -2 // 上一个食物桶的位置
	for i, ch := range hamsters {
		if ch == '.' {
			continue
		}
		// 仓鼠
		if lastFood == i-1 {
			// 左侧有食物桶，嚼嚼嚼
		} else if i < len(hamsters)-1 && hamsters[i+1] == '.' { // 先考虑在右侧放食物桶
			lastFood = i + 1
			ans++
		} else if i > 0 && hamsters[i-1] == '.' { // 再考虑在左侧放食物桶
			ans++ // 无需记录 lastFood（太远了）
		} else { // 饿死了 X_X
			return -1
		}
	}
	return
}
