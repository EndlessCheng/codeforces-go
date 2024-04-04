package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF612E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	ans := make([]int, n+1)
	vis := make([]bool, n+1)
	rt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if vis[a[i]] {
			continue
		}
		sz := 0
		for v := a[i]; !vis[v]; v = a[v] {
			vis[v] = true
			sz++
		}
		if sz&1 > 0 {
			x, y := a[i], a[i]
			for j := sz / 2; j > 0; j-- {
				x = a[x]
			}
			ans[i] = x
			for j := sz - 1; j > 0; j-- {
				ans[y] = a[x]
				x, y = a[x], a[y]
			}
		} else {
			x := rt[sz]
			if x == 0 {
				rt[sz] = i
				continue
			}
			rt[sz] = 0
			y := i
			for j := sz; j > 0; j-- {
				ans[x] = y
				ans[y] = a[x]
				x, y = a[x], a[y]
			}
		}
	}
	for i := 2; i <= n; i += 2 {
		if rt[i] > 0 {
			Fprint(out, -1)
			return
		}
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF612E(os.Stdin, os.Stdout) }
