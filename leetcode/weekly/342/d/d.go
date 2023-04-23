package main

// https://space.bilibili.com/206214
func minOperations(nums []int) int {
	n, gcdAll, cnt1 := len(nums), 0, 0
	for _, x := range nums {
		gcdAll = gcd(gcdAll, x)
		if x == 1 {
			cnt1++
		}
	}
	if gcdAll > 1 {
		return -1
	}
	if cnt1 > 0 {
		return n - cnt1
	}

	minSize := n
	type result struct{ gcd, i int }
	a := []result{}
	for i, x := range nums {
		for j, r := range a {
			a[j].gcd = gcd(r.gcd, x)
		}
		a = append(a, result{x, i})

		// 去重
		j := 0
		for _, q := range a[1:] {
			if a[j].gcd != q.gcd {
				j++
				a[j] = q
			} else {
				a[j].i = q.i
			}
		}
		a = a[:j+1]

		if a[0].gcd == 1 {
			minSize = min(minSize, i-a[0].i)
		}
	}
	return minSize + n - 1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
