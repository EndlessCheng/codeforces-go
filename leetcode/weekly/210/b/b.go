package main

// github.com/EndlessCheng/codeforces-go
func maximalNetworkRank(n int, es [][]int) (ans int) {
	d := make([]int, n)
	conn := make([][]bool, n)
	for i := range conn {
		conn[i] = make([]bool, n)
	}
	for _, e := range es {
		v, w := e[0], e[1]
		d[v]++
		d[w]++
		conn[v][w] = true
		conn[w][v] = true
	}
	for i, v := range d {
		for j := i + 1; j < n; j++ {
			s := v + d[j]
			if conn[i][j] {
				s--
			}
			if s > ans {
				ans = s
			}
		}
	}
	return
}
