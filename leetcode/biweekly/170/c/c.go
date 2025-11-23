package main

// https://space.bilibili.com/206214
func lexSmallestNegatedPerm(n int, target int64) []int {
	t := int(target)
	mx := n * (n + 1) / 2
	if t > mx || -t > mx || (mx-t)%2 != 0 {
		return nil
	}
	negS := (mx - t) / 2 // 取负号的元素（的绝对值）之和

	ans := make([]int, n)
	l, r := 0, n-1
	// 从 1,2,...,n 中选一些数，元素和等于 neg
	// 为了让取反后字典序尽量小，从大往小选
	for x := n; x > 0; x-- {
		if negS >= x {
			negS -= x
			ans[l] = -x
			l++
		} else {
			// 大的正数填在末尾
			ans[r] = x
			r--
		}
	}

	return ans
}
