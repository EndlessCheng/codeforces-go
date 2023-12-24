package main

// https://space.bilibili.com/206214
func incremovableSubarrayCount(a []int) int64 {
	n := len(a)
	i := 0
	for i < n-1 && a[i] < a[i+1] {
		i++
	}
	if i == n-1 { // 每个非空子数组都可以移除
		return int64(n) * int64(n+1) / 2
	}

	ans := int64(i + 2)
	for j := n - 1; j > 0 && (j == n-1 || a[j] < a[j+1]); j-- {
		for i >= 0 && a[i] >= a[j] {
			i--
		}
		ans += int64(i + 2)
	}
	return ans
}
