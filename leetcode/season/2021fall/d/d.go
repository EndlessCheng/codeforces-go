package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func circleGame(toys [][]int, circles [][]int, r0 int) (ans int) {
	sort.Slice(circles, func(i, j int) bool { a, b := circles[i], circles[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })

	// 将横坐标相同的圈分为一组
	type pair struct {
		x  int
		ys []int
	}
	a, y := []pair{}, -1
	for _, p := range circles {
		if len(a) == 0 || p[0] > a[len(a)-1].x {
			a = append(a, pair{p[0], []int{p[1]}})
			y = -1
		} else if p[1] > y { // 去重
			a[len(a)-1].ys = append(a[len(a)-1].ys, p[1])
			y = p[1]
		}
	}

	for _, t := range toys {
		x, y, r := t[0], t[1], t[2]
		if r > r0 {
			continue
		}
		i := sort.Search(len(a), func(i int) bool { return a[i].x+r0 >= x+r })
		for ; i < len(a) && a[i].x-r0 <= x-r; i++ {
			cx, ys := a[i].x, a[i].ys
			j := sort.SearchInts(ys, y)
			if j < len(ys) {
				if cy := ys[j]; (x-cx)*(x-cx)+(y-cy)*(y-cy) <= (r0-r)*(r0-r) {
					ans++
					break
				}
			}
			if j > 0 {
				if cy := ys[j-1]; (x-cx)*(x-cx)+(y-cy)*(y-cy) <= (r0-r)*(r0-r) {
					ans++
					break
				}
			}
		}
	}
	return
}
