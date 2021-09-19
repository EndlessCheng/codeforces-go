package main

// 前缀最大值+后缀最小值

// github.com/EndlessCheng/codeforces-go
func sumOfBeauties(a []int) (ans int) {
	n := len(a)
	sufMin := make([]int, n) // 后缀最小值
	sufMin[n-1] = a[n-1]
	for i := n - 2; i > 1; i-- {
		sufMin[i] = min(sufMin[i+1], a[i])
	}
	preMax := a[0] // 前缀最大值
	for i := 1; i < n-1; i++ {
		v := a[i]
		if preMax < v && v < sufMin[i+1] {
			ans += 2
		} else if a[i-1] < v && v < a[i+1] {
			ans++
		}
		if v > preMax {
			preMax = v
		}
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
