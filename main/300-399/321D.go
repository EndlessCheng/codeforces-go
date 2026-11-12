package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf321D(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n int
	Fscan(in, &n)
	a := make([][]int, n+1)
	for i := range a {
		a[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			Fscan(in, &a[i][j])
		}
	}

	b := make([]int, n+1)
	c := make([]int, n+1)
	ans := 0
	m := (n + 1) / 2
	for i := range 1 << m {
		for j := range m {
			b[j+1] = i >> j & 1
		}
		for j := 1; j < m; j++ {
			b[j+m] = b[j] ^ b[m]
		}
		res := 0
		for j := 1; j <= n; j++ {
			if b[j] != 0 {
				res -= a[j][m]
			} else {
				res += a[j][m]
			}
		}
		for j := 1; j < m; j++ {
			s1, s2 := 0, 0
			for k := 1; k <= n; k++ {
				if b[k] != 0 {
					c[k] = a[k][j] - a[k][j+m]
				} else {
					c[k] = a[k][j] + a[k][j+m]
				}
			}
			for k := 1; k < m; k++ {
				s1 += abs(c[k] + c[k+m])
				s2 += abs(c[k] - c[k+m])
			}
			s1 += c[m]
			s2 -= c[m]
			res += max(s1, s2)
		}
		ans = max(ans, res)
	}
	Fprint(out, ans)
}

//func main() { cf321D(bufio.NewReader(os.Stdin), os.Stdout) }
