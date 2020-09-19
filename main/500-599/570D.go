package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF570D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v, t, d int
	Fscan(in, &n, &q)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &v)
		g[v-1] = append(g[v-1], w)
	}
	var s string
	Fscan(in, &s)
	tin := make([]int, n)
	tout := make([]int, n)
	ts := make([][]int, n+1)
	sum := make([][]int, n+1)
	for i := range sum {
		sum[i] = []int{0}
	}
	var f func(v, d int)
	f = func(v, d int) {
		t++
		tin[v] = t
		ts[d] = append(ts[d], t) // 记录当前深度的时间戳
		sum[d] = append(sum[d], sum[d][len(sum[d])-1]^1<<(s[v]-'a')) // 记录当前深度的异或前缀和
		for _, w := range g[v] {
			f(w, d+1)
		}
		tout[v] = t
	}
	f(0, 1)

	for ; q > 0; q-- {
		Fscan(in, &v, &d)
		v--
		// 根据出入时间戳确定范围
		l := sort.SearchInts(ts[d], tin[v])
		r := sort.SearchInts(ts[d], tout[v]+1)
		if s := sum[d][r] ^ sum[d][l]; s&(s-1) == 0 { // 二进制至多一个 1
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { CF570D(os.Stdin, os.Stdout) }
