package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p1725(in io.Reader, out io.Writer) {
	var n, l, r int
	Fscan(in, &n, &l, &r)
	a := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := make([]int, n+1)
	for i := 1; i < l; i++ {
		f[i] = -1e18
	}
	q := []int{0}
	j := 1
	for i := l; i <= n; i++ {
		for j <= i-l {
			for len(q) > 0 && f[j] >= f[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, j)
			j++
		}
		if q[0] < i-r {
			q = q[1:]
		}
		f[i] = f[q[0]] + a[i]
	}
	Fprint(out, slices.Max(f[n-r+1:]))
}

//func main() { p1725(bufio.NewReader(os.Stdin), os.Stdout) }
