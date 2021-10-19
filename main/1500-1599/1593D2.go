package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1593D2(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		c := map[int]int{}
		for i := range a {
			Fscan(in, &a[i])
			c[a[i]]++
		}
		for _, c := range c {
			if c >= n/2 {
				Fprintln(out, -1)
				continue o
			}
		}

		sort.Ints(a)
		ans := 0
		f := func(base, d int) bool {
			c := 0
			for _, v := range a {
				if abs(v-base)%d == 0 {
					c++
				}
			}
			return c >= n/2
		}
		for i, v := range a {
			for _, w := range a[:i] {
				for d, m := 1, v-w; d*d <= m; d++ {
					if m%d == 0 {
						if f(w, m/d) {
							ans = max(ans, m/d)
							break
						}
						if f(w, d) {
							ans = max(ans, d)
						}
					}
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1593D2(os.Stdin, os.Stdout) }
