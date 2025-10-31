package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2167D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for x := 2; ; x++ {
			for _, v := range a {
				if gcd(x, v) == 1 {
					Fprintln(out, x)
					continue o
				}
			}
		}
	}
}

//func main() { cf2167D(bufio.NewReader(os.Stdin), os.Stdout) }
