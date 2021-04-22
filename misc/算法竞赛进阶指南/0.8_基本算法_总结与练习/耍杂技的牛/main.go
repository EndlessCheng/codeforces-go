package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://www.luogu.com.cn/problem/P1842

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, s int
	Fscan(in, &n)
	a := make([]struct{ w, s int }, n)
	for i := range a {
		Fscan(in, &a[i].w, &a[i].s)
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.w+a.s < b.w+b.s })

	ans := int(-1e18)
	for _, p := range a {
		ans = max(ans, s-p.s)
		s += p.w
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
