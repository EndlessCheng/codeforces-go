package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1916B(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int) int { return a / gcd(a, b) * b }
	var T, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b)
		if b%a > 0 {
			Fprintln(out, lcm(a, b))
		} else {
			Fprintln(out, b/a*b)
		}
	}
}

//func main() { cf1916B(bufio.NewReader(os.Stdin), os.Stdout) }
