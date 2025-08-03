package main

import (
	. "fmt"
	"io"
	"math"
	"slices"
)

// https://github.com/EndlessCheng
func cf1656F(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		suf := 0
		for i := range a {
			Fscan(in, &a[i])
			suf += a[i]
		}

		slices.Sort(a)
		if a[0]*(n-2)+suf > 0 || a[n-1]*(n-2)+suf < 0 {
			Fprintln(out, "INF")
			continue
		}

		ans := math.MinInt
		pre := 0
		suf -= a[n-1]
		for i, v := range a {
			f := a[0]*suf - v*(a[0]*(n-i-1)+suf) + a[n-1]*pre - v*(a[n-1]*i+pre)
			ans = max(ans, f)
			pre += v
			suf -= v
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1656F(bufio.NewReader(os.Stdin), os.Stdout) }
