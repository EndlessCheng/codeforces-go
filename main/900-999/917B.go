package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF917B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w byte
	var m int
	var s []byte
	Fscan(in, &n, &m)
	g := make([][]byte, n)
	for i := range g {
		g[i] = make([]byte, n)
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &s)
		g[v-1][w-1] = s[0] & 31
	}

	dp := make([][][27]byte, n)
	for i := range dp {
		dp[i] = make([][27]byte, n)
	}
	var f func(x, y, c byte) byte
	f = func(x, y, c byte) (r byte) {
		d := &dp[x][y][c]
		if *d > 0 {
			return *d
		}
		defer func() { *d = r }()
		for i, b := range g[x] {
			if b > 0 && b >= c && f(y, byte(i), b) == 'B' {
				return 'A'
			}
		}
		return 'B'
	}
	for i := byte(0); i < n; i++ {
		for j := byte(0); j < n; j++ {
			Fprintf(out, "%c", f(i, j, 0))
		}
		Fprintln(out)
	}
}

//func main() { CF917B(os.Stdin, os.Stdout) }
