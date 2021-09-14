package main

// github.com/EndlessCheng/codeforces-go
func unhappyFriends(n int, preferences [][]int, pairs [][]int) (ans int) {
	a := make([][]int, n)
	for i, row := range preferences {
		a[i] = make([]int, n)
		for j, v := range row {
			a[i][v] = n - j
		}
	}
	f := func(x, y, u, v int) bool { return a[x][u] > a[x][y] && a[u][x] > a[u][v] }
	unh := make([]bool, n)
	f2 := func(x, y, u, v int) { unh[x] = unh[x] || f(x, y, u, v) || f(x, y, v, u) }
	for i, p := range pairs {
		for j, q := range pairs {
			if j != i {
				f2(p[0], p[1], q[0], q[1])
				f2(p[1], p[0], q[0], q[1])
			}
		}
	}
	for _, b := range unh {
		if b {
			ans++
		}
	}
	return
}
