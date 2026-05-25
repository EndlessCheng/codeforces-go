package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2154E(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.SortFunc(a, func(a, b int) int { return b - a })

		s := make([]int, n+1)
		for i, v := range a {
			s[i+1] = s[i] + v
		}

		f := func(i, d int) int {
			r := max(i+1, n-d*k)
			return s[i-d] + s[r] - s[i] + (n-r+d)*a[i]
		}

		ans := 0
		for i := range n {
			l, r := 0, i+1
			for l+1 < r {
				mid := (l + r) / 2
				if f(i, mid) >= f(i, mid-1) {
					l = mid
				} else {
					r = mid
				}
			}
			ans = max(ans, f(i, l))
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2154E(bufio.NewReader(os.Stdin), os.Stdout) }
