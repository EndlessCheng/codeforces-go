package main

import "fmt"

// github.com/EndlessCheng/codeforces-go
func numberOfRounds(startTime, finishTime string) int {
	var h1, m1, h2, m2 int
	fmt.Sscanf(startTime, "%d:%d", &h1, &m1)
	fmt.Sscanf(finishTime, "%d:%d", &h2, &m2)
	if startTime > finishTime { h2 += 24 } // 玩了个通宵
	s, t := h1*60+m1, h2*60+m2
	return (t - t%15 - s - (15-s%15)%15) / 15
}
