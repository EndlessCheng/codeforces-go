package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1467B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if n < 5 {
			Fprintln(out, 0)
			continue
		}

		f := func(i int) bool { return a[i-1] < a[i] && a[i] > a[i+1] || a[i-1] > a[i] && a[i] < a[i+1] }
		cnt := 0
		b := make([]bool, n)
		for i := 1; i < n-1; i++ {
			if f(i) {
				cnt++
				b[i] = true
			}
		}

		ans := cnt
		for i := 1; i < n-1; i++ {
			if !b[i] {
				continue
			}
			if i == 1 || i == n-2 {
				c := cnt - 1
				if i == 1 && b[2] {
					c--
				} else if i == n-2 && b[n-3] {
					c--
				}
				if c < ans {
					ans = c
				}
			} else {
				tmp := a[i]
				g := func(v int) {
					a[i] = v
					c := cnt - 1
					if b[i-1] && !f(i-1) {
						c--
					} else if !b[i-1] && f(i-1) {
						c++
					}
					if b[i+1] && !f(i+1) {
						c--
					} else if !b[i+1] && f(i+1) {
						c++
					}
					if c < ans {
						ans = c
					}
				}
				g(a[i-1])
				g(a[i+1])
				a[i] = tmp
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1467B(os.Stdin, os.Stdout) }
