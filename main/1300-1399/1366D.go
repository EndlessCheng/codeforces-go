package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1366D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 1e7
	lpf := make([]int, mx+1)
	lpf[1] = 1
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	var n, v int
	Fscan(in, &n)
	a := make([]interface{}, n)
	b := make([]interface{}, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		p := lpf[v]
		for v /= p; lpf[v] == p; v /= p {
		}
		if v > 1 {
			a[i] = p
			b[i] = v
		} else {
			a[i] = -1
			b[i] = -1
		}
	}
	Fprintln(out, a...)
	Fprintln(out, b...)
}

//func main() { CF1366D(os.Stdin, os.Stdout) }
