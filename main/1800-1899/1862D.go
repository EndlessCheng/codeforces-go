package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf1862D(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		x := int(math.Sqrt(float64(n*8 + 1)))
		if x*x > n*8+1 {
			x--
		}
		x = (1 + x) / 2
		Fprintln(out, x+n-x*(x-1)/2)
	}
}

//func main() { cf1862D(bufio.NewReader(os.Stdin), os.Stdout) }
