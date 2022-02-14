package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1637D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans, maxW := 0, 0
		a := make([][2]int, n)
		for i := range a {
			Fscan(in, &a[i][0])
			ans += a[i][0] * a[i][0]
			maxW += a[i][0]
		}
		for i := range a {
			Fscan(in, &a[i][1])
			ans += a[i][1] * a[i][1]
			maxW += a[i][1]
		}
		if n == 1 {
			Fprintln(out, 0)
			continue
		}
		dp := make([]bool, maxW+1)
		dp[0] = true
		for _, g := range a {
		next:
			for j := maxW; j >= 0; j-- {
				for _, w := range g {
					if w <= j && dp[j-w] {
						dp[j] = true
						continue next
					}
				}
				dp[j] = false
			}
		}
		ans *= n - 2
		mi, miv := int(1e9), 0
		for v, b := range dp {
			if b {
				d := abs(maxW - 2*v)
				if d < mi {
					mi, miv = d, v
				}
			}
		}
		Fprintln(out, ans+miv*miv+(maxW-miv)*(maxW-miv))
	}
}

//func main() { CF1637D(os.Stdin, os.Stdout) }
