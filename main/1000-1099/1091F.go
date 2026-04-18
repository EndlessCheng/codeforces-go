package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1091F(in io.Reader, out io.Writer) {
	var n, ans, g, w int
	var s string
	Fscan(in, &n)
	l := make([]int, n)
	for i := range l {
		Fscan(in, &l[i])
		l[i] *= 2
		ans += l[i]
	}
	Fscan(in, &s)

	mn := 5
	for i, b := range s {
		if b == 'G' {
			g += l[i]
		} else if b == 'W' {
			w += l[i]
			mn = 3
		}

		t := l[i]
		e := min(t/2, w)
		ans += e * 2
		t -= e * 2
		w -= e

		e = min(t/2, g)
		ans += e * 4
		t -= e * 2
		g -= e

		ans += mn * t
	}
	Fprint(out, ans/2)
}

//func main() { cf1091F(bufio.NewReader(os.Stdin), os.Stdout) }
