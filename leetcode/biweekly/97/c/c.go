package main

// https://space.bilibili.com/206214
func maximizeWin(prizePositions []int, k int) (ans int) {
	n := len(prizePositions)
	if k*2+1 >= prizePositions[n-1]-prizePositions[0] {
		return n
	}
	mx, left, right := 0, 0, 0
	for mid, p := range prizePositions {
		// 把 prizePositions[mid] 视作第二条线段的左端点，计算第二条线段可以覆盖的最大奖品下标
		for right < n && prizePositions[right]-p <= k {
			right++
		}
		// 循环结束后，right-1 是第二条线段可以覆盖的最大奖品下标
		ans = max(ans, mx+right-mid)
		// 把 prizePositions[mid] 视作第一条线段的右端点，计算第一条线段可以覆盖的最小奖品下标
		for p-prizePositions[left] > k {
			left++
		}
		// 循环结束后，left 是第一条线段可以覆盖的最小奖品下标
		mx = max(mx, mid-left+1)
	}
	return
}

func maximizeWin2(prizePositions []int, k int) (ans int) {
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
