package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1638C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		p := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			a[i]--
			p[a[i]] = i
		}
		ans := 0
		for i, min := n-1, n; i >= 0; i-- {
			if a[p[i]] < 0 {
				continue
			}
			if a[p[i]] < min {
				ans++
			}
			for j := p[i]; j < n && a[j] >= 0; j++ {
				if a[j] < min {
					min = a[j]
				}
				a[j] = -1
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1638C(os.Stdin, os.Stdout) }
