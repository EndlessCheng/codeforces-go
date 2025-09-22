package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf1473B(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		g := gcd(len(s), len(t))
		x := strings.Repeat(s, len(t)/g)
		if x == strings.Repeat(t, len(s)/g) {
			Fprintln(out, x)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { cf1473B(bufio.NewReader(os.Stdin), os.Stdout) }
