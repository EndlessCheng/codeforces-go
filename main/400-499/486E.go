package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF486E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	pre := make([]int, n)
	g := []int{}
	for i := range a {
		Fscan(in, &a[i])
		v := a[i]
		p := sort.SearchInts(g, v)
		if p < len(g) {
			g[p] = v
		} else {
			g = append(g, v)
		}
		pre[i] = p + 1
	}

	suf := make([]int, n)
	g = g[:0]
	for i := n - 1; i >= 0; i-- {
		v := -a[i]
		p := sort.SearchInts(g, v)
		if p < len(g) {
			g[p] = v
		} else {
			g = append(g, v)
		}
		suf[i] = p + 1
	}

	lis := len(g)
	ans := make([]byte, n)
	cnt := make([]int, n+1)
	for i, p := range pre {
		if p+suf[i]-1 == lis {
			ans[i] = '3' // 暂定是 3
			cnt[p]++     // 如果有多个相同的 p，则 ans[i] = '2'（下面判断）
		} else {
			ans[i] = '1' // 不在任何 LIS 上
		}
	}

	for i, tp := range ans {
		if tp == '3' && cnt[pre[i]] > 1 {
			ans[i] = '2'
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF486E(bufio.NewReader(os.Stdin), os.Stdout) }
