package main

// https://space.bilibili.com/206214
func minArrivalsToDiscard(arrivals []int, w, m int) (ans int) {
	cnt := map[int]int{}
	for i, x := range arrivals {
		// x 进入窗口
		if cnt[x] == m { // x 的个数已达上限
			// 注意 x 在未来要离开窗口，但由于已经丢弃，不能计入
			// 这里直接置为 0，未来离开窗口就是 cnt[0]--，不影响答案
			arrivals[i] = 0
			ans++
		} else {
			cnt[x]++
		}

		// 左端点元素离开窗口，为下一个循环做准备
		left := i + 1 - w
		if left >= 0 {
			cnt[arrivals[left]]--
		}
	}
	return
}
