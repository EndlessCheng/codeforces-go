package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF689D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n int
	Fscan(in, &n)
	st := make([][18][2]int, n)
	for i := range st {
		Fscan(in, &st[i][0][0])
	}
	for i := range st {
		Fscan(in, &st[i][0][1])
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j][0] = max(st[i][j-1][0], st[i+1<<(j-1)][j-1][0])
			st[i][j][1] = min(st[i][j-1][1], st[i+1<<(j-1)][j-1][1])
		}
	}
	query := func(l, r int) int {
		k := bits.Len(uint(r-l)) - 1
		p, q := st[l][k], st[r-1<<k][k]
		return max(p[0], q[0]) - min(p[1], q[1])
	}

	ans := int64(0)
	for i, l, r := 1, 0, 0; i <= n; i++ {
		for l < i && query(l, i) > 0 {
			l++
		}
		for r < i && query(r, i) >= 0 {
			r++
		}
		ans += int64(r - l)
	}
	Fprint(out, ans)
}

//func main() { CF689D(os.Stdin, os.Stdout) }
