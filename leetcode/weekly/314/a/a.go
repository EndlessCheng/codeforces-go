package main

// https://space.bilibili.com/206214
func hardestWorker(n int, logs [][]int) int {
	ans, maxT := logs[0][0], logs[0][1]
	for i := 1; i < len(logs); i++ {
		t := logs[i][1] - logs[i-1][1]
		if t > maxT || t == maxT && logs[i][0] < ans {
			ans, maxT = logs[i][0], t
		}
	}
	return ans
}
