package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF786C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, clock int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	time := make([]int, n+1)
	f := func(k int) (res int) {
		clock++
		left := 0
		for _, v := range a {
			if time[v] == clock {
				continue
			}
			if left == 0 {
				res++
				clock++
				left = k
			}
			time[v] = clock
			left--
		}
		return
	}
	ans := make([]int, n+1)
	var solve func(l, r int)
	solve = func(l, r int) {
		if l > r {
			return
		}
		r1, r2 := f(l), f(r)
		if r1 == r2 {
			for i := l; i <= r; i++ {
				ans[i] = r1
			}
			return
		}
		ans[l] = r1
		ans[r] = r2
		mid := (l + r) / 2
		solve(l+1, mid)
		solve(mid+1, r-1)
	}
	solve(1, n)
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF786C(os.Stdin, os.Stdout) }
