package main

// https://space.bilibili.com/206214
func countOfPairs(n, x, y int) []int64 {
	if x > y {
		x, y = y, x
	}

	ans := make([]int64, n)
	if x+1 >= y {
		for i := 1; i < n; i++ {
			ans[i-1] = int64(n-i) * 2
		}
		return ans
	}

	diff := make([]int, n+1)
	add := func(l, r int) {
		diff[l]++
		diff[r+1]--
	}

	for i := 1; i < n; i++ {
		if i <= x {
			k := (x + y + 1) / 2
			add(1, k-i)
			add(x-i+2, x-i+y-k)
			add(x-i+1, x-i+1+n-y)
		} else if i < (x+y)/2 {
			k := i + (y-x+1)/2
			add(1, k-i)
			add(i-x+2, i-x+y-k)
			add(i-x+1, i-x+1+n-y)
		} else {
			add(1, n-i)
		}
	}

	sumD := int64(0)
	for i, d := range diff[1:] {
		sumD += int64(d)
		ans[i] = sumD * 2
	}
	return ans
}
