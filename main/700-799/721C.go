package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF721C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, maxT int
	Fscan(in, &n, &m, &maxT)
	es := make([][3]int, m)
	for i := range es {
		Fscan(in, &es[i][0], &es[i][1], &es[i][2])
	}

	const mx = 5001
	dp := make([][mx]int, n+1) // int32
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = maxT + 1
		}
	}
	dp[1][1] = 0
	fa := make([][mx]int16, n+1)

	// 不需要拓扑，直接跑 n-1 次就行了
	ans := 0
	for i := 2; i <= n; i++ {
		for _, e := range es {
			v, w, t := e[0], e[1], e[2]
			if sumT := dp[i-1][v] + t; sumT < dp[i][w] {
				dp[i][w] = sumT
				fa[i][w] = int16(v)
			}
		}
		if dp[i][n] <= maxT {
			ans = i
		}
	}
	Fprintln(out, ans)
	path := make([]interface{}, ans)
	o := int16(n)
	for i := ans; i > 0; i-- {
		path[i-1] = o
		o = fa[i][o]
	}
	Fprint(out, path...)
}

//func main() { CF721C(os.Stdin, os.Stdout) }
