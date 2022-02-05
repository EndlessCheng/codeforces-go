package main

import (
	"fmt"
	"math"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func minCostSetTime(startAt, moveCost, pushCost, sec int) int {
	ans := math.MaxInt32
	calc := func(s string) {
		cost := pushCost * len(s)
		cur := startAt
		for _, ch := range s {
			if int(ch&15) != cur {
				cost += moveCost
				cur = int(ch & 15)
			}
		}
		if cost < ans { ans = cost }
	}
	if sec >= 60 && sec < 6000 {
		calc(fmt.Sprintf("%d%02d", sec/60, sec%60))
	}
	if sec < 100 {
		calc(strconv.Itoa(sec)) // 仅输入秒数
	} else if sec%60 < 40 {
		calc(fmt.Sprintf("%d%d", sec/60-1, sec%60+60)) // 借一分钟给秒数
	}
	return ans
}
