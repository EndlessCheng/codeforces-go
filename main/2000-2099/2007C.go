package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2007C(in io.Reader, out io.Writer) {
	var T, n, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y)
		g := gcd07(x, y)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			a[i] %= g
		}
		slices.Sort(a)

		mx := a[0] - a[n-1] + g
		for i := 1; i < n; i++ {
			mx = max(mx, a[i]-a[i-1])
		}
		Fprintln(out, g-mx)
	}
}

//func main() { cf2007C(bufio.NewReader(os.Stdin), os.Stdout) }
func gcd07(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
