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
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	type pair struct{ v, i int }
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(x, y int) (res int) {
		if x+y == n {
			return
		}
		d := &dp[x][y]
		if *d >= 0 {
			return *d
		}
		defer func() { *d = res }()
		p := a[x+y]
		f1 := f(x+1, y) + p.v*abs(x-p.i)
		f2 := f(x, y+1) + p.v*abs(n-1-y-p.i)
		if f1 > f2 {
			return f1
		}
		return f2
	}
	Fprint(_w, f(0, 0))
}

func main() { run(os.Stdin, os.Stdout) }
