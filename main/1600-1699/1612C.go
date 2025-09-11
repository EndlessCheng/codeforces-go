package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func ceilSqrt(x int) int {
	return int(math.Ceil(math.Sqrt(float64(x))))
}

func isqrt(x int) int {
	rt := int(math.Sqrt(float64(x)))
	if rt*rt > x {
		rt--
	}
	return rt
}

func cf1612C(in io.Reader, out io.Writer) {
	var T, k, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k, &x)
		k2 := k*2 - 1
		if k*(k+1)/2 >= x {
			Fprintln(out, ceilSqrt(x*8+1)/2)
		} else if k*k >= x {
			x -= k * (k + 1) / 2
			Fprintln(out, k+(k2-isqrt(k2*k2-x*8)+1)/2)
		} else {
			Fprintln(out, k2)
		}
	}
}

//func main() { cf1612C(bufio.NewReader(os.Stdin), os.Stdout) }
