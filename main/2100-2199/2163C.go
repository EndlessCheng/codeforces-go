package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2163C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}

		type pair struct{ min, max int }
		suf := make([]pair, n+1)
		suf[n].min = n * 2
		for i := n - 1; i >= 0; i-- {
			suf[i].min = min(suf[i+1].min, b[i])
			suf[i].max = max(suf[i+1].max, b[i])
		}

		r2l := make([]int, n*2+1)
		preMax, preMin := 0, n*2
		for i, v := range a {
			preMax = max(preMax, v)
			preMin = min(preMin, v)
			l := min(preMin, suf[i].min)
			r := max(preMax, suf[i].max)
			r2l[r] = max(r2l[r], l)
		}

		ans, maxL := 0, 0
		for _, l := range r2l {
			maxL = max(maxL, l)
			ans += maxL
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2163C(bufio.NewReader(os.Stdin), os.Stdout) }
