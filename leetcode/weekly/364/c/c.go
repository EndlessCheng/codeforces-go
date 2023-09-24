package main

// https://space.bilibili.com/206214
func maximumSumOfHeights(a []int) int64 {
	ans := 0
	n := len(a)
	suf := make([]int, n+1)
	st := []int{n} // 哨兵
	sum := 0
	for i := n - 1; i >= 0; i-- {
		x := a[i]
		for len(st) > 1 && x <= a[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			sum -= a[j] * (st[len(st)-1] - j) // 撤销之前加到 sum 中的
		}
		sum += x * (st[len(st)-1] - i) // 从 i 到 st[len(st)-1]-1 都是 x
		suf[i] = sum
		st = append(st, i)
	}
	ans = sum

	st = []int{-1} // 哨兵
	pre := 0
	for i, x := range a {
		for len(st) > 1 && x <= a[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			pre -= a[j] * (j - st[len(st)-1]) // 撤销之前加到 pre 中的
		}
		pre += x * (i - st[len(st)-1]) // 从 st[len(st)-1]+1 到 i 都是 x
		ans = max(ans, pre+suf[i+1])
		st = append(st, i)
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
