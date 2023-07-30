package main

// https://space.bilibili.com/206214
func numberOfEmployeesWhoMetTarget(hours []int, target int) (ans int) {
	for _, h := range hours {
		if h >= target {
			ans++
		}
	}
	return
}
