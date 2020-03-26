package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1311D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var t, a, b, c, ti, tj, tk int
	for Fscan(in, &t); t > 0; t-- {
		ans := int(1e9)
		Fscan(in, &a, &b, &c)
		for i := 1; i < 2*a; i++ {
			for j := i; j < 2*b; j += i {
				k := c - c%j
				if k == 0 {
					k = j
				} else if abs(k+j-c) < abs(k-c) {
					k += j
				}
				if v := abs(i-a) + abs(j-b) + abs(k-c); v < ans {
					ans, ti, tj, tk = v, i, j, k
				}
			}
		}
		Fprintln(out, ans)
		Fprintln(out, ti, tj, tk)
	}
}

//func main() { CF1311D(os.Stdin, os.Stdout) }
