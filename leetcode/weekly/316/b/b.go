package main

// https://space.bilibili.com/206214
func subarrayGCD(nums []int, k int) (ans int) {
	type result struct{ v, i int }
	var a []result
	i0 := -1
	for i, v := range nums {
		if v%k > 0 {
			a = nil
			i0 = i
			continue
		}
		for j, p := range a {
			a[j].v = gcd(p.v, v)
		}
		a = append(a, result{v, i})
		j := 0
		for _, q := range a[1:] {
			if a[j].v != q.v {
				j++
				a[j] = q
			} else {
				a[j].i = q.i
			}
		}
		a = a[:j+1]
		if a[0].v == k {
			ans += a[0].i - i0
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
