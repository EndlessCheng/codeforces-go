package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2124C(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int) int { return a / gcd(a, b) * b }
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g, l := 0, 1
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			g = gcd(g, v)
			l = lcm(l, v/g)
		}
		Fprintln(out, l)
	}
}

//func main() { cf2124C(bufio.NewReader(os.Stdin), os.Stdout) }
