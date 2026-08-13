package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1463F(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var c, x, y int
	Fscan(in, &c, &x, &y)
	var f func(int) int
	f = func(n int) int {
		g := gcd(x, y)
		p := x + y
		x /= g
		y /= g
		if g > 1 {
			return f(n/g+1)*(n%g) + f(n/g)*(g-n%g)
		}
		ans := 0
		d := [2]int{0, -n}
		for j := range 2 {
			for i := range p {
				ex := 0
				if i*x%p < n%p {
					ex = 1
				}
				d = [2]int{max(d[0], d[1]), d[0] + n/p + ex}
			}
			ans = max(ans, d[j])
			d = [2]int{-n, 0}
		}
		return ans
	}
	Fprint(out, f(c))
}

//func main() { cf1463F(bufio.NewReader(os.Stdin), os.Stdout) }
