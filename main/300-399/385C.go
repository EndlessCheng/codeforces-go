package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF385C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 1e7
	lpf := make([]int, mx+1)
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	var n, v, q, l, r int
	c := make([]int, mx+1)
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		for v > 1 {
			p := lpf[v]
			for v /= p; lpf[v] == p; v /= p {
			}
			c[p]++
		}
	}
	for i := 2; i < mx; i++ {
		c[i+1] += c[i]
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		if l > mx {
			l = mx
		}
		if r > mx {
			r = mx
		}
		Fprintln(out, c[r]-c[l-1])
	}
}

//func main() { CF385C(os.Stdin, os.Stdout) }
