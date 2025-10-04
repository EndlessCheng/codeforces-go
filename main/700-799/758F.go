package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf758F(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n, l, r, ans int
	Fscan(in, &n, &l, &r)
	if n == 1 {
		Fprint(out, r-l+1)
	} else if n == 2 {
		Fprint(out, (r-l+1)*(r-l))
	} else {
		for p := 1; math.Pow(float64(p), float64(n-1)) <= float64(r); p++ {
			powP := math.Pow(float64(p), float64(n-1))
			for q := 1; q < p; q++ {
				if gcd(p, q) == 1 {
					powQ := math.Pow(float64(q), float64(n-1))
					ans += max(r/int(powP)-(l-1)/int(powQ), 0)
				}
			}
		}
		Fprint(out, ans*2)
	}
}

//func main() { cf758F(os.Stdin, os.Stdout) }
