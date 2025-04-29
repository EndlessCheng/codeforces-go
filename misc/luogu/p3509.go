package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3509(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, m int
	Fscan(in, &n, &k, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	f := make([]int, n)
	r := k
	for i, v := range a {
		for r+1 < n && a[r+1]-v < v-a[r-k] {
			r++
		}
		if v-a[r-k] >= a[r]-v {
			f[i] = r - k
		} else {
			f[i] = r
		}
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = i
	}
	for ; m > 0; m /= 2 {
		if m%2 > 0 {
			for i, j := range ans {
				ans[i] = f[j]
			}
		}
		nf := make([]int, n)
		for i, to := range f {
			nf[i] = f[to]
		}
		f = nf
	}
	for _, v := range ans {
		Fprint(out, v+1, " ")
	}
}

//func main() { p3509(bufio.NewReader(os.Stdin), os.Stdout) }
