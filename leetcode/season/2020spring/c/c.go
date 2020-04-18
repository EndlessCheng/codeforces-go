package main

import "sort"

func getTriggerTime(increase [][]int, req [][]int) (ans []int) {
	n := len(req)
	ans = make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	increase = append(increase, []int{0, 0, 0})
	for i := range req {
		req[i] = append(req[i], i, 0)
	}

	for cur := 0; cur < 3; cur++ {
		sort.Slice(req, func(i, j int) bool { return req[i][cur] < req[j][cur] })
		val, j := 0, 0
		for i, inc := range increase {
			for ; j < n && req[j][cur] <= val; j++ {
				if idx := req[j][3]; ans[idx] < i {
					ans[idx] = i
				}
				req[j][4]++
			}
			val += inc[cur]
		}
	}
	for _, r := range req {
		if r[4] < 3 {
			ans[r[3]] = -1
		}
	}
	return
}
