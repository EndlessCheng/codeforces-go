package main

import (
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	ans := 999
	c := make([]int, n)
	d := make([]int, n)
	var f func(int, bool)
	f = func(i int, odd bool) {
		if i == n {
			if odd {
				return
			}
			c := append([]int{}, c...)
			sort.Ints(c)
			d := append([]int{}, d...)
			res := 0
		o:
			for i, v := range c {
				for j := i; j < n; j++ {
					w := d[j]
					if v != w && v != -w || w < 0 != ((i-j)&1 > 0) {
						continue
					}
					res += j - i
					for k := j - 1; k >= i; k-- {
						d[k] = -d[k]
						d[k], d[k+1] = d[k+1], d[k]
					}
					continue o
				}
				return
			}
			if res < ans {
				ans = res
			}
			return
		}
		c[i] = a[i]
		d[i] = a[i]
		f(i+1, odd)
		c[i] = b[i]
		d[i] = -b[i]
		f(i+1, !odd)
	}
	f(0, false)
	if ans == 999 {
		ans = -1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
