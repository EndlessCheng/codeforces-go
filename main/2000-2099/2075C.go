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
			a[i] = min(a[i], u-1)
			s += a[i]
		}
		slices.Sort(a)

		ans := 0
		l, r := 0, n-1
		for l < r {
			if a[l]+a[r] < u {
				s -= a[l]
				l++
			} else {
				s -= a[r]
				ans += s - (r-l)*(u-1-a[r])
				r--
			}
		}
		Fprintln(out, ans*2)
	}
}

//func main() { cf2075C(bufio.NewReader(os.Stdin), os.Stdout) }
