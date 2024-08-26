package main

// https://space.bilibili.com/206214
func maximizeWin(prizePositions []int, k int) (ans int) {
	n := len(prizePositions)
	if k*2+1 >= prizePositions[n-1]-prizePositions[0] {
		return n
	}
	pre := make([]int, n+1)
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
