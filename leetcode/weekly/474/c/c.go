package main

import "sort"

// https://space.bilibili.com/206214
func minimumTime1(d, r []int) int64 {
	d1, d2 := d[0], d[1]
	r1, r2 := r[0], r[1]
	l := lcm(r1, r2)

	// 库函数是左闭右开区间
	left := d1 + d2
	right := (d1+d2)*2 - 1
	ans := left + sort.Search(right-left, func(t int) bool {
		t += left
		return d1 <= t-t/r1 && d2 <= t-t/r2 && d1+d2 <= t-t/l
	})
	return int64(ans)
}

func f(d, r int) int {
	return (d-1)/(r-1) + 1
}

func minimumTime(d, r []int) int64 {
	d1, d2 := d[0], d[1]
	r1, r2 := r[0], r[1]
	l := lcm(r1, r2)
	return int64(max(f(d1, r1), f(d2, r2), f(d1+d2, l)))
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
