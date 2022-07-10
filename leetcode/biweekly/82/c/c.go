package main

import "sort"

// https://space.bilibili.com/206214/dynamic
func minSumSquareDiff(a, nums2 []int, k1, k2 int) int64 {
	ans, sum := 0, 0
	for i, v := range a {
		a[i] = abs(v - nums2[i])
		sum += a[i]
		ans += a[i] * a[i]
	}
	k := k1 + k2
	if sum <= k {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	a = append(a, 0) // 哨兵
	for i, v := range a {
		i++
		ans -= v * v
		if c := i * (v - a[i]); c < k {
			k -= c
			continue
		}
		v -= k / i
		ans += k%i*(v-1)*(v-1) + (i-k%i)*v*v
		break
	}
	return int64(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
