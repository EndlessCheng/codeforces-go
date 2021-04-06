package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1503C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ a, c int }

	var n, mxAC int
	Fscan(in, &n)
	a := make([]pair, n)
	ans := int64(0)
	for i := range a {
		Fscan(in, &a[i].a, &a[i].c)
		ans += int64(a[i].c) // 基础花费
	}
	sort.Slice(a, func(i, j int) bool { return a[i].a < a[j].a })
	for i, p := range a[:n-1] {
		if p.a+p.c > mxAC {
			mxAC = p.a + p.c
		}
		if d := a[i+1].a - mxAC; d > 0 { // 从 i+1 之前的任何位置都无法白嫖到 i+1
			ans += int64(d)
		}
	}
	Fprint(out, ans)
}

//func main() { CF1503C(os.Stdin, os.Stdout) }
