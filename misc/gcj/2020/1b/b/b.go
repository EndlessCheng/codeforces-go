package main

import . "fmt"

func main() {
	var done bool
	search := func(l, r int, f func(int) bool) int {
		for !done && l < r {
			m := (l + r) >> 1
			if f(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	var s []byte
	io := func(x, y int) byte {
		Println(x, y)
		Scan(&s)
		if s[0] == 'C' {
			done = true
		}
		return s[0]
	}

	var t, a, b, x, y int
	for Scan(&t, &a, &b); t > 0; t-- {
		done = false

		// 枚举 9 个位置找一个在圆内的点
	o:
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				x, y = 1e9/2*i, 1e9/2*j
				if io(x, y) != 'M' {
					break o
				}
			}
		}

		// 二分查找过点 (x,y) 的水平弦和竖直弦
		xl := search(-1e9, x, func(xx int) bool { return io(xx, y) == 'H' })
		xr := search(x, 1e9+1, func(xx int) bool { return io(xx, y) == 'M' })
		yl := search(-1e9, y, func(yy int) bool { return io(x, yy) == 'H' })
		yr := search(y, 1e9+1, func(yy int) bool { return io(x, yy) == 'M' })

		// 由于弦的两端不一定在圆上，先确定一个大致的圆心，然后暴力 check 周围的点
		ox, oy := (xl+xr)/2, (yl+yr)/2
		for i := ox - 5; i <= ox+5; i++ {
			for j := oy - 5; j <= oy+5; j++ {
				if !done {
					io(i, j)
				}
			}
		}
	}
}
