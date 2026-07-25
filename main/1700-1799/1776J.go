package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1776J(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	c := make([]int, n)
	for i := range c {
		Fscan(in, &c[i])
	}
	dis := make([][]int, n*2)
	for i := range dis {
		dis[i] = make([]int, n*2)
		for j := range dis[i] {
			dis[i][j] = 1e9
		}
		dis[i][i] = 0
	}
	for i := 0; i < n; i++ {
		dis[i][i+n] = k
		dis[i+n][i] = k
	}

	for range m {
		var x, y int
		Fscan(in, &x, &y)
		x--
		y--
		if c[x] != c[y] {
			dis[x][y+n] = 1
			dis[y+n][x] = 1
			dis[x+n][y] = 1
			dis[y][x+n] = 1
		} else {
			dis[x][y] = 1
			dis[y][x] = 1
			dis[x+n][y+n] = 1
			dis[y+n][x+n] = 1
		}
	}

	for k := range dis {
		for i := range dis {
			for j := range dis {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, (dis[i][j]+dis[i][j+n]+k)>>1)
		}
	}
	Fprint(out, ans)
}

//func main() { cf1776J(bufio.NewReader(os.Stdin), os.Stdout) }
