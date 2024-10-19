package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1941F(in io.Reader, out io.Writer) {
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}
		c := make([]int, k)
		for i := range c {
			Fscan(in, &c[i])
		}

		ans, se, mxI := a[1]-a[0], 0, 1
		for i := 2; i < n; i++ {
			d := a[i] - a[i-1]
			if d > ans {
				ans, se, mxI = d, ans, i
			} else if d > se {
				se = d
			}
		}
		mid := (a[mxI-1] + a[mxI]) / 2

		slices.Sort(b)
		slices.Sort(c)
		j := k - 1
		for _, v := range b {
			for j >= 0 && v+c[j] > mid {
				j--
			}
			if j >= 0 {
				ans = min(ans, a[mxI]-v-c[j])
			}
			if j+1 < k {
				ans = min(ans, v+c[j+1]-a[mxI-1])
			}
		}
		Fprintln(out, max(ans, se))
	}
}

//func main() { cf1941F(bufio.NewReader(os.Stdin), os.Stdout) }
