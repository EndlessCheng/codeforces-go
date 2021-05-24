package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1004D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var t, v, mx int
	Fscan(in, &t)
	cnt := make([]int, t)
	for range cnt {
		Fscan(in, &v)
		cnt[v]++
		if v > mx {
			mx = v
		}
	}
	row := 0
	for ; row+1 < t && row*4+4 == cnt[row+1]; row++ {
	}
	f := func(n, m int) bool {
		col := n - 1 - row + m - 1 - mx
		if row >= n || col < 0 || col >= m {
			return false
		}
		c2 := make([]int, t)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				v := abs(i-row) + abs(j-col)
				if v >= t {
					return false
				}
				if c2[v]++; c2[v] > cnt[v] {
					return false
				}
			}
		}
		for i, c := range c2 {
			if c != cnt[i] {
				return false
			}
		}
		Fprintln(out, n, m)
		Fprintln(out, row+1, col+1)
		return true
	}
	for d := 1; d*d <= t; d++ {
		if t%d == 0 && (f(d, t/d) || f(t/d, d)) {
			return
		}
	}
	Fprint(out, -1)
}

//func main() { CF1004D(os.Stdin, os.Stdout) }
