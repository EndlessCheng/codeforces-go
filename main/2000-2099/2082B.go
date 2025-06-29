package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2082B(in io.Reader, out io.Writer) {
	var T, x, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &n, &m)
		v := x
		for m := m; v > 1 && m > 0; m-- {
			v = (v + 1) / 2
		}
		mn := v >> n

		v = x >> n
		for ; v > 1 && m > 0; m-- {
			v = (v + 1) / 2
		}

		Fprintln(out, mn, v)
	}
}

//func main() { cf2082B(bufio.NewReader(os.Stdin), os.Stdout) }
