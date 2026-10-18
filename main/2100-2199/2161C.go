package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2161C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s += a[i]
		}
		k := s / x

		slices.Sort(a)
		ans := 0
		for _, v := range a[n-k:] {
			ans += v
		}
		Fprintln(out, ans)

		s = 0
		i, j := 0, n-k
		for range n {
			if j < n && s%x+a[j] >= x {
				s += a[j]
				Fprint(out, a[j], " ")
				j++
			} else {
				s += a[i]
				Fprint(out, a[i], " ")
				i++
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2161C(bufio.NewReader(os.Stdin), os.Stdout) }
