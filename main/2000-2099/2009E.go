package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func isqrt9(x uint) int {
	rt := uint(math.Sqrt(float64(x)))
	if rt*rt > x {
		rt--
	}
	return int(rt)
}

func cf2009E(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		s := (k*2 + n - 1) * n / 2
		b := k*2 - 1
		i := (isqrt9(uint(b*b)+uint(s)*4) - b) / 2
		f := func(i int) int { return (k*2+i-1)*i - s }
		Fprintln(out, min(-f(i), f(i+1)))
	}
}

//func main() { cf2009E(bufio.NewReader(os.Stdin), os.Stdout) }
