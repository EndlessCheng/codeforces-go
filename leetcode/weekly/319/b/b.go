package main

// https://space.bilibili.com/206214
func subarrayLCM(nums []int, k int) (ans int) {
	type result struct{ lcm, i int }
	var a []result
	i0 := -1
	for i, x := range nums {
		if k%x > 0 {
			a = nil
			i0 = i
			continue
		}
		for j, p := range a {
			a[j].lcm = p.lcm / gcd(p.lcm, x) * x
		}
		for len(a) > 0 && k%a[0].lcm > 0 { // 去除不合法的 LCM
			a = a[1:]
		}
		a = append(a, result{x, i})
		j := 0
		for _, q := range a[1:] {
			if a[j].lcm != q.lcm {
				j++
				a[j] = q
			} else {
				a[j].i = q.i
			}
		}
		a = a[:j+1]
		if a[0].lcm == k {
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
