package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// https://github.com/EndlessCheng
func isqrt(x int) int {
	rt := int(math.Sqrt(float64(x)))
	if rt*rt > x {
		rt--
	}
	return rt
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, c, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &c, &d)
		ans := 0
		for p10 := 1; p10 <= c+d; p10 *= 10 {
			l := c*p10*10 + max(c+1, p10)
			r := c*p10*10 + min(c+d, p10*10-1)
			if l <= r {
				ans += isqrt(r) - isqrt(l-1)
			}
		}
		Fprintln(out, ans)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
