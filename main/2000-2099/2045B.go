package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2045B(in io.Reader, out io.Writer) {
	var n, d, s int
	Fscan(in, &n, &d, &s)
	if d < s {
		Fprint(out, s)
		return
	}
	n /= s
	d /= s
	t := min(n, d*2)
	f := func(i int) bool {
		return t/i <= d && t-t/i <= d
	}
	for i := 1; i*i <= t; i++ {
		if t%i == 0 && (f(i) || f(t/i)) {
			Fprint(out, t*s)
			return
		}
	}
	Fprint(out, (t-1)*s)
}

//func main() { cf2045B(os.Stdin, os.Stdout) }
