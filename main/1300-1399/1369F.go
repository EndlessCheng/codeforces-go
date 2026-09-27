package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1369F(in io.Reader, out io.Writer) {
	b2i := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}
	var f func(int, int) int
	f = func(s, e int) int {
		if e%2 != 0 || s > e/4*2 {
			return (e - s) & 1
		}
		return b2i(s > e/4) | f(s, e/4)
	}

	var n int
	Fscan(in, &n)
	l, w := 1, 0
	for n > 0 && w != l {
		n--
		var x, y int
		Fscan(in, &x, &y)
		l = b2i(x > y/2 || f(x, y/2) != 0) ^ w
		w ^= f(x, y)
	}
	Fprintln(out, w, l)
}

//func main() { cf1369F(bufio.NewReader(os.Stdin), os.Stdout) }
