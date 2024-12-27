package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2051D(in io.Reader, out io.Writer) {
	var T, n, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s += a[i]
		}
		lower, upper := s-y, s-x
		slices.Sort(a)
		ans := 0
		left, right := n, n
		for j, x := range a {
			for right > 0 && a[right-1] > upper-x {
				right--
			}
			for left > 0 && a[left-1] >= lower-x {
				left--
			}
			ans += min(right, j) - min(left, j)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2051D(bufio.NewReader(os.Stdin), os.Stdout) }
