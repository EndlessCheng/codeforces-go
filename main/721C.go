package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol721C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m, maxT int
	Fscan(in, &n, &m, &maxT)
	edges := make([][3]int, m)
	for i := range edges {
		e := &edges[i]
		Fscan(in, &e[0], &e[1], &e[2])
	}

	const mx = 5001
	dp := [mx][mx]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = maxT + 1
		}
	}
	dp[1][1] = 0
	fa := [mx][mx]int{}

	ans := 0
	for i := 2; i <= n; i++ {
		for _, e := range edges {
			v, w, t := e[0], e[1], e[2]
			if sumT := dp[i-1][v] + t; sumT < dp[i][w] {
				dp[i][w] = sumT
				fa[i][w] = v
			}
		}
		if dp[i][n] <= maxT {
			ans = i
		}
	}
	Fprintln(out, ans)
	path := make([]interface{}, ans)
	o := n
	for i := ans; i > 0; i-- {
		path[i-1] = o
		o = fa[i][o]
	}
	Fprint(out, path...)
}

//func main() {
//	Sol721C(os.Stdin, os.Stdout)
//}
