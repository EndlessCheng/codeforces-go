package main

// https://space.bilibili.com/206214
func maxFreeTime(eventTime, k int, startTime, endTime []int) (ans int) {
	n := len(startTime)
	free := make([]int, n+1)
	free[0] = startTime[0]
	for i := 1; i < n; i++ {
		free[i] = startTime[i] - endTime[i-1]
	}
	free[n] = eventTime - endTime[n-1]

	s := 0
	for i, f := range free {
		s += f
		if i < k {
			continue
		}
		ans = max(ans, s)
		s -= free[i-k]
	}
	return
}

func maxFreeTime2(eventTime, k int, startTime, endTime []int) (ans int) {
	n := len(startTime)
	get := func(i int) int {
		if i == 0 {
			return startTime[0]
		}
		if i == n {
			return eventTime - endTime[n-1]
		}
		return startTime[i] - endTime[i-1]
	}

	s := 0
	for i := range n + 1 {
		s += get(i)
		if i < k {
			continue
		}
		ans = max(ans, s)
		s -= get(i - k)
	}
	return
}
