package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2075C(in io.Reader, out io.Writer) {
	var T, u, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &u, &n)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s += min(a[i]+1, u)
		}
		slices.Sort(a)

		ans := 0
		l, r := 0, n-1
		for l < r {
			if a[l]+a[r] < u {
				s -= a[l] + 1
				l++
			} else {
				s -= min(a[r]+1, u)
				ans += s - (r-l)*max(u-a[r], 1)
				r--
			}
		}
		Fprintln(out, ans*2)
	}
}

//func main() { cf2075C(bufio.NewReader(os.Stdin), os.Stdout) }
