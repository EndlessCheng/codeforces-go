package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1789F(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	n := len(s)
	ans := uint8(0)

	const N = 80
	f := [N][N]uint8{}
	for m := 1; m < n; m++ {
		for i, x := range s[:m] {
			for j, y := range s[m:] {
				if x == y {
					f[i+1][j+1] = f[i][j] + 1
				} else {
					f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
				}
			}
		}
		ans = max(ans, f[m][n-m]*2)
	}

	if n >= 5 {
		nxt := make([][26]int, n+1)
		for j := range nxt[n] {
			nxt[n][j] = n
		}
		for i := n - 1; i >= 0; i-- {
			nxt[i] = nxt[i+1]
			s[i] -= 'a'
			nxt[i][s[i]] = i
		}
		match := func(t []byte) uint8 {
			i := -1
			for j := 0; ; j++ {
				i = nxt[i+1][t[j%len(t)]]
				if i == n {
					if j < len(t)*5 {
						return 0
					}
					return uint8(j - j%len(t))
				}
			}
		}

		_t := [N]byte{}
		st := 0
		for i := range 5 {
			sz := n / 5
			if i < n%5 {
				sz++
			}
			curS := s[st : st+sz]
			st += sz
			for sub := 1; sub < 1<<len(curS); sub++ {
				t := _t[:0]
				for j := uint(sub); j > 0; j &= j - 1 {
					t = append(t, curS[bits.TrailingZeros(j)])
				}
				ans = max(ans, match(t))
			}
		}
	}

	g := [N][N][N]uint8{}
	low := int(ans/3 + 1)
	for m1 := low; m1 < n; m1++ {
		for m2 := m1 + low; m2 < n; m2++ {
			for i, x := range s[:m1] {
				for j, y := range s[m1:m2] {
					for k, z := range s[m2:] {
						if x == y && y == z {
							g[i+1][j+1][k+1] = g[i][j][k] + 1
						} else {
							g[i+1][j+1][k+1] = max(g[i][j+1][k+1], g[i+1][j][k+1], g[i+1][j+1][k])
						}
					}
				}
			}
			ans = max(ans, g[m1][m2-m1][n-m2]*3)
		}
	}

	Fprint(out, ans)
}

//func main() { cf1789F(os.Stdin, os.Stdout) }
