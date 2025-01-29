package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1909B(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		or, and := 0, -1
		for range n {
			Fscan(in, &v)
			or |= v
			and &= v
		}
		or ^= and
		Fprintln(out, or&-or<<1)
	}
}

//func main() { cf1909B(bufio.NewReader(os.Stdin), os.Stdout) }
