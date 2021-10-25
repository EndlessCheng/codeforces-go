package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1512F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &c)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n-1)
		for i := range b {
			Fscan(in, &b[i])
		}

		ans := int(2e9)
		cur, day := 0, 0
		for i, v := range a {
			if cur >= c {
				break
			}
			if i < n-1 {
				res := (c - cur + v - 1) / v
				ans = min(ans, res+day)
				if b[i] >= c {
					break
				}
				if cur < b[i] {
					d := (b[i] - cur + v - 1) / v
					day += d + 1
					cur += d*v - b[i]
				} else {
					cur -= b[i]
					day++
				}
			} else {
				res := (c - cur + v - 1) / v
				ans = min(ans, res+day)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1512F(os.Stdin, os.Stdout) }
