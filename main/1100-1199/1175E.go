package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1175E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const mx = 19

	var n, q, l, r int
	Fscan(in, &n, &q)
	f := make([][mx]int, 5e5+1)
	for ; n > 0; n-- {
		Fscan(in, &l, &r)
		f[l][0] = max(f[l][0], r)
	}
	for i := 1; i < len(f); i++ {
		f[i][0] = max(f[i][0], f[i-1][0])
	}
	for i := 1; i < mx; i++ {
		for p := range f {
			f[p][i] = f[f[p][i-1]][i-1]
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		ans := 0
		for i := mx - 1; i >= 0; i-- {
			if f[l][i] < r {
				l = f[l][i]
				ans |= 1 << i
			}
		}
		if f[l][0] >= r {
			Fprintln(out, ans+1)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { CF1175E(os.Stdin, os.Stdout) }
