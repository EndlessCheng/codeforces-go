package main

// https://space.bilibili.com/206214
func goodIndices(a []int, k int) (ans []int) {
n := len(a)
	dec := make([]int,n+1)
	dec[n-1] = 1
	for i := len(a) - 2; i >= 0; i-- {
		v := a[i]
		if v <=a[i+1] {
			dec[i] = dec[i+1]+1
		} else {
			dec[i] = 1
		}
	}

	pos := make([]int, n)
	pos[0] = 1
	for i := 1; i < n; i++ {
		v := a[i]
		if v <=a[i-1] {
			pos[i] = pos[i-1]+1
		} else {
			pos[i] = 1
		}
	}

	for i := 1; i < n-1; i++ {
		if pos[i-1] >=k && dec[i+1] >= k {
			ans = append(ans, i)
		}
	}

	return
}
