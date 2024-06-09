package main

// https://space.bilibili.com/206214
func findWinningPlayer(skills []int, k int) (mxI int) {
	win := 0
	for i := 1; i < len(skills) && win < k; i++ {
		if skills[i] > skills[mxI] { // 新的最大值
			mxI = i
			win = 0
		}
		win++ // 获胜回合 +1
	}
	return
}
