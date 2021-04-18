package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1301B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		f := func(k int) (m int) {
			for i := 1; i < n; i++ {
				v := a[i]
				if v < 0 {
					v = k
				}
				w := a[i-1]
				if w < 0 {
					w = k
				}
				m = max(m, abs(v-w))
			}
			return
		}
		k := sort.Search(1e9, func(m int) bool { return f(m) < f(m+1) })
		Fprintln(out, f(k), k)
	}
}

//func main() { CF1301B(os.Stdin, os.Stdout) }
