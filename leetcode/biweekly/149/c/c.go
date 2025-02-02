package main

// https://space.bilibili.com/206214
func maxFreeTime(eventTime int, startTime, endTime []int) (ans int) {
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

	a, b, c := 0, -1, -1
	for i := 1; i <= n; i++ {
		sz := get(i)
		if sz > get(a) {
			a, b, c = i, a, b
		} else if b < 0 || sz > get(b) {
			b, c = i, b
		} else if c < 0 || sz > get(c) {
			c = i
		}
	}

	for i, e := range endTime {
		sz := e - startTime[i]
		if i != a && i+1 != a && sz <= get(a) ||
			i != b && i+1 != b && sz <= get(b) ||
			sz <= get(c) {
			ans = max(ans, get(i)+sz+get(i+1))
		} else {
			ans = max(ans, get(i)+get(i+1))
		}
	}
	return ans
}
