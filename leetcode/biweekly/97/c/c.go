package main

// https://space.bilibili.com/206214
func maximizeWin(prizePositions []int, k int) (ans int) {
	pre := make([]int, len(prizePositions)+1)
	left := 0
	for right, p := range prizePositions {
		for p-prizePositions[left] > k {
			left++
		}
		ans = max(ans, right-left+1+pre[left])
		pre[right+1] = max(pre[right], right-left+1)
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
