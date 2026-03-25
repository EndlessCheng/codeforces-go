package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2060G(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n)
		cnt := 0
		for i := range b {
			Fscan(in, &b[i])
			if a[i] > b[i] {
				cnt++
				a[i], b[i] = b[i], a[i]
			}
		}

		p := make([]int, n)
		for i := range p {
			p[i] = i
		}
		slices.SortFunc(p, func(i, j int) int { return a[i] - a[j] })

		ok := true
		pre := 0
		for i := 1; i < n; i++ {
			ok = ok && b[p[i-1]] < b[p[i]]
			if b[p[i-1]] < a[p[i]] {
				if (i-pre)%2 > 0 {
					cnt = 0
				}
				pre = i
			}
		}
		if n%2 > 0 {
			cnt = 0
		}

		if ok && cnt%2 == 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf2060G(bufio.NewReader(os.Stdin), os.Stdout) }
