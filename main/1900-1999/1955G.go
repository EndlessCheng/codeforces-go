package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1955G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 1_000_001
	D := [mx][]int{}
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i {
			D[j] = append(D[j], i)
		}
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]int, n)
		vis := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
			vis[i] = make([]int, m)
		}
		ds := D[gcd(a[0][0], a[n-1][m-1])]
		for i := len(ds) - 1; ; i-- {
			d := ds[i]
			var dfs func(int, int) bool
			dfs = func(x, y int) bool {
				if x == 0 && y == 0 {
					return true
				}
				vis[x][y] = d
				return y > 0 && vis[x][y-1] != d && a[x][y-1]%d == 0 && dfs(x, y-1) ||
					x > 0 && vis[x-1][y] != d && a[x-1][y]%d == 0 && dfs(x-1, y)
			}
			if dfs(n-1, m-1) {
				Fprintln(out, d)
				break
			}
		}
	}
}

//func main() { cf1955G(os.Stdin, os.Stdout) }
