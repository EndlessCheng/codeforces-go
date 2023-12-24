package main

// https://space.bilibili.com/206214
func incremovableSubarrayCount(a []int) int {
	n := len(a)
	i := 0
	for i < n-1 && a[i] < a[i+1] {
		i++
	}
	if i == n-1 { // 每个非空子数组都可以移除
		return n * (n + 1) / 2
	}

	ans := i + 2
	for j := n - 1; j == n-1 || a[j] < a[j+1]; j-- {
		for i >= 0 && a[i] >= a[j] {
			i--
		}
		ans += i + 2
	}
	return ans
}
