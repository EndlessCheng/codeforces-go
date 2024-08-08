package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf1902C(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if n == 1 {
			Fprintln(out, 1)
			continue
		}

		slices.Sort(a)
		g := 0
		for i := 1; i < n; i++ {
			g = gcd(g, a[i]-a[i-1])
		}

		s := 0
		for _, v := range a {
			s += (a[n-1] - v) / g
		}

		ex := n
		for i := n - 1; i > 0; i-- {
			if a[i]-a[i-1] > g {
				ex = n - i
				break
			}
		}
		Fprintln(out, s+ex)
	}
}

//func main() { cf1902C(bufio.NewReader(os.Stdin), os.Stdout) }
