package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func transpose45(a [][]byte) [][]byte {
	n, m := len(a), len(a[0])
	b := make([][]byte, m)
	for i := range b {
		b[i] = make([]byte, n)
		for j, r := range a {
			b[i][j] = r[i]
		}
	}
	return b
}

func cf845F(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	if n < m {
		a = transpose45(a)
		n, m = m, n
	}

	walls := make([]int, n)
	for i, row := range a {
		for j, b := range row {
			if b == 'x' {
				walls[i] |= 1 << j
			}
		}
	}

	dp := make([][][][2][2]int, n)
	for i := range dp {
		dp[i] = make([][][2][2]int, m)
		for j := range dp[i] {
			dp[i][j] = make([][2][2]int, 1<<m)
			for k := range dp[i][j] {
				dp[i][j][k] = [2][2]int{{-1, -1}, {-1, -1}}
			}
		}
	}
	var f func(int, int, int, int, int) int
	f = func(i, j, down, right, miss int) (res int) {
		if i == n {
			return 1
		}
		if j == m {
			return f(i+1, 0, down, 0, miss)
		}
		down &^= walls[i]
		p := &dp[i][j][down][right][miss]
		if *p >= 0 {
			return *p
		}
		defer func() { *p = res }()
		if walls[i]>>j&1 > 0 {
			return f(i, j+1, down, 0, miss)
		}
		res = f(i, j+1, down|1<<j, 1, miss)
		if right > 0 || down>>j&1 > 0 {
			res = (res + f(i, j+1, down, right, miss)) % mod
		} else if miss == 0 {
			res = (res + f(i, j+1, down, right, 1)) % mod
		}
		return
	}
	Fprint(out, f(0, 0, 0, 0, 0))
}

//func main() { cf845F(bufio.NewReader(os.Stdin), os.Stdout) }
