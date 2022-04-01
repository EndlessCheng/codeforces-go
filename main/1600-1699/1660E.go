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
		n *= 2
		sum := make([][]int, n+1)
		sum[0] = make([]int, n+1)
		for i, r := range a {
			sum[i+1] = make([]int, n+1)
			for j, v := range r {
				sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + int(v&1)
			}
		}

		ans := int(1e9)
		for s := n/2 + 1; s < n/2*3; s++ {
			l := max(0, n-s)
			r := min(n-1, n*2-s-1)
			c1 := 0
			for i, j := s+l-n, l; j <= r; j++ {
				c1 += int(a[i][j] & 1)
				if j >= l+n/2-1 {
					x, y := i-n/2, j-n/2
					if j >= l+n/2 {
						c1 -= int(a[x][y] & 1)
					}
					x++
					y++
					s1 := sum[i+1][j+1] - sum[i+1][y] - sum[x][j+1] + sum[x][y]
					ans = min(ans, s1-2*c1+n/2)
				}
				i++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1660E(os.Stdin, os.Stdout) }
