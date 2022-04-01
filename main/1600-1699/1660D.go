package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1660D(_r io.Reader, _w io.Writer) {
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
		mx, ll, rr := 0, 0, 0
		for i := 0; i < n; i++ {
			if a[i] == 0 {
				continue
			}

			st := i
			c, neg, l, r := 0, false, -1, 0
			for ; i < n && a[i] != 0; i++ {
				if a[i] == -2 || a[i] == 2 {
					c++
				}
				if a[i] < 0 {
					neg = !neg
					if l < 0 {
						l = i
					}
					r = i
				}
			}

			if neg {
				cl, cr := 0, 0
				for j := st; j < i; j++ {
					if a[j] == -2 || a[j] == 2 {
						if j < r {
							cl++
						}
						if j > l {
							cr++
						}
					}
				}
				if cl > cr {
					c, l, r = cl, st, r-1
				} else {
					c, l, r = cr, l+1, i-1
				}
			} else {
				l, r = st, i-1
			}

			if c > mx {
				mx, ll, rr = c, l, r
			}
		}

		if mx == 0 {
			Fprintln(out, n, 0)
		} else {
			Fprintln(out, ll, n-rr-1)
		}
	}
}

//func main() { CF1660D(os.Stdin, os.Stdout) }
