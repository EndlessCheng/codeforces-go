package main

import (
	"sort"
)

/* 二分+暴力枚举

注意到半径很小，我们可以枚举每个玩具，并暴力枚举该玩具**周围**是否有可以套住该玩具的圈。

具体来说，将 $\textit{circles}$ 排序后，将横坐标相同的圈分为一组。对每个玩具，可以套住该玩具的圈，在横坐标上，必然满足圈的最右端点不小于玩具的最右端点，且圈的最左端点不超过玩具的最左端点。对于纵坐标也是类似。这样我们可以二分要枚举的圈的初始位置，并向后遍历即可。

*/

// github.com/EndlessCheng/codeforces-go
func circleGame(toys [][]int, circles [][]int, r0 int) (ans int) {
	sort.Slice(circles, func(i, j int) bool { a, b := circles[i], circles[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })

	// 将横坐标相同的圈分为一组
	type pair struct{ x, y int }
	a, y := [][]pair{}, -1
	for _, p := range circles {
		if len(a) == 0 || p[0] > a[len(a)-1][0].x {
			a = append(a, []pair{{p[0], p[1]}})
			y = -1
		} else if p[1] > y { // 去重
			a[len(a)-1] = append(a[len(a)-1], pair{p[0], p[1]})
			y = p[1]
		}
	}

outer:
	for _, t := range toys {
		x, y, r := t[0], t[1], t[2]
		if r > r0 {
			continue
		}
		i := sort.Search(len(a), func(i int) bool { return a[i][0].x+r0 >= x+r })
		for ; i < len(a); i++ {
			col := a[i]
			if col[0].x-r0 > x-r {
				break
			}
			j := sort.Search(len(col), func(j int) bool { return col[j].y+r0 >= y+r })
			for ; j < len(col); j++ {
				c := col[j]
				if c.y-r0 > y-r {
					break
				}
				if (x-c.x)*(x-c.x)+(y-c.y)*(y-c.y) <= (r0-r)*(r0-r) {
					ans++
					continue outer
				}
			}
		}
	}
	return
}
