package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1660E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
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

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]string, n)
		for i := range a {
			Fscan(in, &a[i])
			a[i] += a[i]
		}
		a = append(a, a...)
		m := n * 2
		s := make([][]int, m+1)
		s[0] = make([]int, m+1)
		for i, r := range a {
			s[i+1] = make([]int, m+1)
			for j, v := range r {
				s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j] + int(v&1)
			}
		}

		ans := int(1e9)
		for k := n + 1; k < n*3; k++ {
			l, r := max(0, m-k), min(m-1, m*2-k-1)
			for i, j, c1 := k+l-m, l, 0; j <= r; j++ {
				c1 += int(a[i][j] & 1)
				if j >= l+n-1 {
					x, y := i-n, j-n
					if j >= l+n {
						c1 -= int(a[x][y] & 1)
					}
					ans = min(ans, s[i+1][j+1]-s[i+1][y+1]-s[x+1][j+1]+s[x+1][y+1]-2*c1+n)
				}
				i++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1660E(os.Stdin, os.Stdout) }
