package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n,s int
	Fscan(in, &n)
	a := make([]struct{ x, y,z int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y,&a[i].z)
		s += a[i].z
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, s*2+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, j int) (res int) {
		if i == n {
			if j > s {
				return 0
			}
			return 1e18
		}
		dv := &dp[i][j]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		
		p := a[i]
		x,y,z := p.x,p.y,p.z
		if x > y {
			return f(i+1, j+z)		
		} else {
			res = f(i+1, j-z)
			d := (x+y+1)/2 - x
			res = min(res, f(i+1, j+z) + d)
		}
		
		return
	}
	ans := f(0, s)
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}