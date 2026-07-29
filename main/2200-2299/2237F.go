package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2237F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}

		add := 0
		if a[1] != 1 {
			a[1] = 1
			add++
		}
		if a[n] != m {
			a[n] = m
			add++
		}

		f := make([]int, n+m+2)
		for i := range f {
			f[i] = -2e9
		}

		f[n] = 1
		l := f[n-1+m]
		r := 1

		for i := 2; i <= n; i++ {
			f[n-i+1] = r
			f[n-i+a[i]] = max(f[n-i+a[i]], l)
			f[n-i+a[i]]++
			r = max(r, f[n-i+a[i]])
			l = max(l, f[n-i+m])
		}

		Fprintln(out, n-f[m]+add)
	}
}

//func main() { cf2237F(bufio.NewReader(os.Stdin), os.Stdout) }
