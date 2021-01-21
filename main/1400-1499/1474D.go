package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1474D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		check := func(a []int) bool {
			b := append([]int(nil), a...)
			n := len(b)
			for i := 1; i < n; i++ {
				if b[i-1] > b[i] {
					return false
				}
				b[i] -= b[i-1]
			}
			return b[n-1] == 0
		}

		if check(a) {
			Fprintln(out, "YES")
			continue
		}

		a[0], a[1] = a[1], a[0]
		if check(a) {
			Fprintln(out, "YES")
			continue
		}
		a[0], a[1] = a[1], a[0]

		a[n-2], a[n-1] = a[n-1], a[n-2]
		if check(a) {
			Fprintln(out, "YES")
			continue
		}
		a[n-2], a[n-1] = a[n-1], a[n-2]

		b := append([]int(nil), a...)
		pre := make([]int, n)
		pre[0] = b[0]
		for i := 1; i < n; i++ {
			if pre[i-1] == -1 || b[i-1] > b[i] {
				pre[i] = -1
			} else {
				b[i] -= b[i-1]
				pre[i] = b[i]
			}
		}

		b = append([]int(nil), a...)
		suf := make([]int, n)
		suf[n-1] = b[n-1]
		for i := n - 2; i >= 0; i-- {
			if suf[i+1] == -1 || b[i+1] > b[i] {
				suf[i] = -1
			} else {
				b[i] -= b[i+1]
				suf[i] = b[i]
			}
		}

		for i := 1; i+2 < n; i++ {
			if pre[i-1] >= 0 && suf[i+2] >= 0 && check([]int{pre[i-1], a[i+1], a[i], suf[i+2]}) {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1474D(os.Stdin, os.Stdout) }
