package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF864E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ w, mxW, v, i int }

	var n int
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].w, &a[i].mxW, &a[i].v)
		a[i].i = i + 1
	}
	sort.Slice(a, func(i, j int) bool { return a[i].mxW < a[j].mxW })

	m := a[n-1].mxW
	dp := make([]int, m)
	fa := make([][]int, n)
	for i := range fa {
		fa[i] = make([]int, m)
		for j := range fa[i] {
			fa[i][j] = -1
		}
	}
	for i, p := range a {
		w := p.w
		for j := p.mxW - 1; j >= w; j-- {
			if d := dp[j-w] + p.v; d > dp[j] {
				dp[j] = d
				fa[i][j] = j - w
			}
		}
	}
	mxJ := 0
	for j, v := range dp {
		if v > dp[mxJ] {
			mxJ = j
		}
	}
	Fprintln(out, dp[mxJ]) // 由于不是从 m-1 开始枚举的，需要寻找一个最大值
	ans := []int{}
	for i, j := n-1, mxJ; i >= 0; i-- {
		if fa[i][j] != -1 {
			ans = append(ans, a[i].i)
			j = fa[i][j]
		}
	}
	Fprintln(out, len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		Fprint(out, ans[i], " ")
	}
}

//func main() { CF864E(os.Stdin, os.Stdout) }
