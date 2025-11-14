package main

import "slices"

func canMakeArithmeticProgression1(arr []int) bool {
	slices.Sort(arr)
	d := arr[1] - arr[0] // 公差
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != d {
			return false
		}
	}
	return true
}

func canMakeArithmeticProgression2(arr []int) bool {
	mx := slices.Max(arr)
	mn := slices.Min(arr)
	if mn == mx { // 特殊情况：公差为 0 的等差数列
		return true
	}

	if (mx-mn)%(len(arr)-1) > 0 { // 公差必须是整数
		return false
	}
	d := (mx - mn) / (len(arr) - 1) // 公差

	has := map[int]bool{}
	for _, x := range arr {
		has[x] = true
	}
	// 检查 mn, mn+d, mn+2d, ..., mx 每个数是否都存在（注意这一共有 n 个数）
	// mn 和 mx 已经有无需检查
	for x := mn + d; x < mx; x += d {
		if !has[x] {
			return false
		}
	}
	return true
}

func canMakeArithmeticProgression(arr []int) bool {
	mx := slices.Max(arr)
	mn := slices.Min(arr)
	if mn == mx { // 特殊情况：公差为 0 的等差数列
		return true
	}

	n := len(arr)
	if (mx-mn)%(n-1) > 0 { // 公差必须是整数
		return false
	}
	d := (mx - mn) / (n - 1) // 公差

	has := make([]bool, n)
	for _, x := range arr {
		k := (x - mn) / d
		if (x-mn)%d > 0 || has[k] { // k 不是整数或者之前遇到过
			return false
		}
		has[k] = true
	}
	return true
}
