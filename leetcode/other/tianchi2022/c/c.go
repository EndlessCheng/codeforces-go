package main

// https://space.bilibili.com/206214
func arrangeBookshelf(a []int, limit int) (ans []int) {
	left := map[int]int{}
	for _, v := range a {
		left[v]++
	}

	st := []int{0}
	cntSt := map[int]int{}
	for _, x := range a {
		if cntSt[x] == limit {
			left[x]--
			continue
		}
		for {
			top := st[len(st)-1]
			if top <= x || left[top] <= limit {
				break
			}
			cntSt[top]--
			left[top]--
			st = st[:len(st)-1]
		}
		st = append(st, x)
		cntSt[x]++
	}
	ans = st[1:]
	return
}
