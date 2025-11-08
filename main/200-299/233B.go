package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func isqrt(x int) int {
	rt := int(math.Sqrt(float64(x)))
	if rt*rt > x {
		rt--
	}
	return rt
}

func cf233B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	for s := 1; s < 82; s++ {
		d := s*s + n*4
		rt := isqrt(d)
		if rt*rt == d && (rt-s)%2 == 0 {
			x := (rt - s) / 2
			sum := 0
			for t := x; t > 0; t /= 10 {
				sum += t % 10
			}
			if sum == s {
				Fprint(out, x)
				return
			}
		}
	}
	Fprint(out, -1)
}

//func main() { cf233B(bufio.NewReader(os.Stdin), os.Stdout) }
