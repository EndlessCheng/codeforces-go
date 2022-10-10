package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF38E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}
	var n int
	Fscan(in, &n)
	a := make([]struct{ x, c int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].c)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

	f := make([]int64, n)
	f[0] = int64(a[0].c)
	for i := 1; i < n; i++ {
		mi := f[0]
		for j := 0; j < i; j++ {
			mi = min(mi, f[j])
			f[j] += int64(a[i].x-a[j].x)
		}
		f[i] = mi + int64(a[i].c)
	}
	ans := f[0]
	for _, v := range f {
		ans = min(ans, v)
	}
	Fprint(out, ans)
}

//func main() { CF38E(os.Stdin, os.Stdout) }
