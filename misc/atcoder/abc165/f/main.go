package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	r := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b&15)
		}
		return
	}

	n := r()
	a := make([]int, n)
	for i := range a {
		a[i] = r()
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		v, w := r()-1, r()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	ans := make([]int, n)
	dp := []int{}
	var f func(v, fa int)
	f = func(v, fa int) {
		if i := sort.SearchInts(dp, a[v]); i < len(dp) {
			old := dp[i]
			dp[i] = a[v]
			ans[v] = len(dp)
			for _, w := range g[v] {
				if w != fa {
					f(w, v)
				}
			}
			dp[i] = old
		} else {
			dp = append(dp, a[v])
			ans[v] = len(dp)
			for _, w := range g[v] {
				if w != fa {
					f(w, v)
				}
			}
			dp = dp[:len(dp)-1]
		}
	}
	f(0, -1)
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }
