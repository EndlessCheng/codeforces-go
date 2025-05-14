package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf744C(in io.Reader, out io.Writer) {
	var n, sr, sb int
	var c string
	Fscan(in, &n)
	a := make([]struct{ c, r, b int }, n)
	for i := range a {
		Fscan(in, &c, &a[i].r, &a[i].b)
		a[i].c = int(c[0] >> 4 & 1)
		sr += a[i].r
		sb += a[i].b
	}

	dp := make([][121]int, 1<<n)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int, int, int) int
	f = func(s, saveR, rc, bc int) (saveB int) {
		if s == 0 {
			if saveR > 0 {
				return -1e9
			}
			return
		}
		p := &dp[s][saveR]
		if *p != -1 {
			return *p
		}
		saveB = -1e9
		for t := uint(s); t > 0; t &= t - 1 {
			i := bits.TrailingZeros(t)
			r := min(a[i].r, rc)
			if r <= saveR {
				saveB = max(saveB, f(s^1<<i, saveR-r, rc+a[i].c, bc+1-a[i].c)+min(a[i].b, bc))
			}
		}
		*p = saveB
		return
	}

	ans := int(1e9)
	for saveR := range min(n*(n-1)/2, sr) + 1 {
		ans = min(ans, max(sr-saveR, sb-f(1<<n-1, saveR, 0, 0)))
	}
	Fprint(out, ans+n)
}

//func main() { cf744C(bufio.NewReader(os.Stdin), os.Stdout) }
