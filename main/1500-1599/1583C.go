package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1583C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, p, q, l, r int
	var pre, cur string
	Fscan(in, &n, &m, &pre)
	bad := make([]bool, m)
	for ; n > 1; n-- {
		Fscan(in, &cur)
		for j := 1; j < m; j++ {
			if cur[j-1] == 'X' && pre[j] == 'X' {
				bad[j] = true
			}
		}
		pre = cur
	}
	left := make([]int, m)
	for i, b := range bad {
		if b {
			p = i
		}
		left[i] = p
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		if left[r-1] >= l {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { CF1583C(os.Stdin, os.Stdout) }
